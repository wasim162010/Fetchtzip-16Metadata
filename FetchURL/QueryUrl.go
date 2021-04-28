package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	s "strings"
)

type Licence struct {
	Name    string `json:"name,omitempty"`
	Details string `json:"details,omitempty"`
}

type source struct {
	Tools    []string `json:"tools,omitempty"`
	Location string   `json:"location,omitempty"`
}

type MetaData struct {
	Name        string  `json:"name,omitempty"`
	Description string  `json:"description,omitempty"`
	Version     string  `json:"version,omitempty"`
	Licence     Licence `json:"license,omitempty"`
	Homepage    string  `json:"homepage,omitempty"`
	Source      source  `json:"source,omitempty"`
}

// var httpUrl = "https://storage.googleapis.com/tzip-16/emoji-in-metadata.json"
// var sha256Url = "sha256://0x7e99ecf3a4490e3044ccdf319898d77380a2fc20aae36b6e40327d678399d17b/https:%2F%2Fstorage.googleapis.com%2Ftzip-16%2Ftaco-shop-metadata.json"
//ipfsUrl := "ipfs://QmcMUKkhXowQjCPtDVVXyFJd7W9LmC92Gs5kYH1KjEisdj"

/*

The program takes the url as a command line param.

For ex :

go run QueryUrl.go "https://storage.googleapis.com/tzip-16/emoji-in-metadata.json"

go run QueryUrl.go  "sha256://0x7e99ecf3a4490e3044ccdf319898d77380a2fc20aae36b6e40327d678399d17b/https:%2F%2Fstorage.googleapis.com%2Ftzip-16%2Ftaco-shop-metadata.json"

*/

func main() {

	urlObj := os.Args[1]
	fmt.Println("url passed as command line argument ", urlObj)
	u, err := url.Parse(urlObj)
	if err != nil {
		log.Fatalln(err)
	}
	scheme := s.ToLower(u.Scheme)

	switch scheme {
	case "http", "https":
		QueryHttpUrl(urlObj)
	case "sha256":
		QuerySHA256Url(urlObj)
	case "ipfs":
		fmt.Println("To be implemented")
	}
}

func QueryHttpUrl(httpUrl string) (MetaData, error) {
	fmt.Println("QueryHttpUrl")
	var metadata MetaData
	resp, err := http.Get(httpUrl)

	if err != nil {
		fmt.Println("Failed to read the metadata ", err)
		return MetaData{}, err
	}

	if resp.StatusCode == 200 {
		//Reding the response body contents
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {

			fmt.Println("Error while reading response body ", err)
			return MetaData{}, err
		}

		err1 := json.Unmarshal(body, &metadata)
		if err1 != nil {
			fmt.Println("Error while unmarshalling ", err)
			return MetaData{}, err
		}

		fmt.Println("Metadata")
		fmt.Println(metadata)

	} else {
		fmt.Println("Error during get. Response is ", resp)
		return MetaData{}, err
	}
	return metadata, nil

}

func QuerySHA256Url(sha256pUrl string) (MetaData, error) {

	fmt.Println("QuerySHA256Url")
	fetchedURL := s.Split(s.Split(sha256pUrl, "://")[1], "/")[1]
	var metadata MetaData

	if s.Contains(fetchedURL, "%2F%2F") {

		fetchedURL = s.Replace(fetchedURL, "%2F%2F", "//", -1)
		if s.Contains(fetchedURL, "%2F") {
			fetchedURL = s.Replace(fetchedURL, "%2F", "/", -1)
		}
		fmt.Println("fetched http url ", fetchedURL)

	}

	metadata, err := QueryHttpUrl(fetchedURL)

	if err != nil {
		fmt.Println("Failed to read the metadata ", err)
		return MetaData{}, err
	}

	return metadata, nil

}
