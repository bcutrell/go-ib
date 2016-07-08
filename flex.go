package main

import (
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"gopkg.in/matryer/try.v1"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"time"
)

type Config struct {
	Query   string
	Token   string
	BaseUrl string
}

type XmlResponse struct {
	Url           string
	ReferenceCode string
	Status        string
}

type ReportXmlResponse struct {
	ErrorCode    string
	Status       string
	ErrorMessage string
}

func Urlencode(c Config) string {
	params := url.Values{}
	params.Add("t", c.Token)
	params.Add("q", c.Query)
	params.Add("v", "3")

	finalUrl := fmt.Sprintf("%s?%s", c.BaseUrl, params.Encode())
	return finalUrl
}

func flex(fullUrl string) []byte {
	resp, err := http.Get(fullUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

func main() {
	fmt.Println("IB Flex Reader")

	// Get Params
	file, _ := os.Open("config.json")
	decoder := json.NewDecoder(file)
	config := Config{}
	err := decoder.Decode(&config)
	if err != nil {
		fmt.Println("error:", err)
	}

	// Add params to URL
	fullUrl := Urlencode(config)
	fmt.Println(fullUrl)

	// Get initial IB response
	resp := flex(fullUrl)

	// Parse and update config values
	var x XmlResponse
	xml.Unmarshal(resp, &x)

	// Update URL
	config.Query = x.ReferenceCode
	config.BaseUrl = x.Url
	fullUrl = Urlencode(config)

	var fullResp []byte
	reportGenerr := try.Do(func(attempt int) (bool, error) {
		var reportGenerr error

		// Get IB report
		fullResp = flex(fullUrl)
		fmt.Println(string(fullResp))

		var x2 ReportXmlResponse
		xml.Unmarshal(fullResp, &x2)

		// reportGenerr != nil
		fmt.Println(x2)
		if x2.ErrorCode == "1019" {
			reportGenerr = errors.New("can't work with 42")
			time.Sleep(1 * time.Minute) // wait a minute
		} else {
			reportGenerr = nil
		}

		return attempt < 5, reportGenerr // try 5 times
	})

	if reportGenerr != nil {
		fmt.Println("error:", err)
	}
	
	// Write file
	file_err := ioutil.WriteFile("result.csv", fullResp, 0644)
	checkError("File Write Error", file_err)
}

func checkError(message string, err error) {
    if err != nil {
        fmt.Println(message, err)
    }
}
