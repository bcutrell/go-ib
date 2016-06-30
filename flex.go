package main

import (
	"encoding/json"
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

func Urlencode(baseUrl string) string {
	params := url.Values{}
	params.Add("t", Token)
	params.Add("q", QueryId)
	params.Add("v", "3")

	finalUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())
	return finalUrl
}

func flex(fullUrl string) string {
	resp, err := http.Get(fullUrl)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
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
	baseUrl := config.BaseUrl

	// Add params to URL
	fullUrl := Urlencode(baseUrl)
	fmt.Println(fullUrl)

	// Get initial IB response
	resp := flex(fullUrl)

	// Write CSV
}
