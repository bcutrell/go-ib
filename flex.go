package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
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
	config.Query = x.ReferenceCode
	config.BaseUrl = x.Url

	// Get IB report
	resp = flex(fullUrl)
	fmt.Println(string(resp))

	// Write CSV
}
