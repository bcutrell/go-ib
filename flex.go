package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func Urlencode(baseUrl string) {
	params := url.Values{}
	params.Add("query", "keyword")

	finalUrl := fmt.Sprintf("%s?%s", baseUrl, params.Encode())
	fmt.Println(finalUrl)
}

func flex() {
	// Params
	// query_id, token, verbose, url
	url := "https://gdcdyn.interactivebrokers.com/Universal/servlet/FlexStatementService.SendRequest"

	// Http request
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(resp)
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	// Return
	// csv file

}

func main() {
	fmt.Println("IB Flex Reader")
	// flex()
	Urlencode("http://theringer.com")
}

// https://blog.gopheracademy.com/advent-2014/reading-config-files-the-go-way/
