package main

import (
	"crypto/sha256"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
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
	//fmt.Println("url passed as command line argument ", urlObj)
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
	var metadata MetaData

	u, err := url.Parse(sha256pUrl)
	if err != nil {
		log.Fatal(err)
	}
	//  url.Path will fetch  http url : "/https://storage.googleapis.com/tzip-16/taco-shop-metadata.json"
	fetchedURL := u.Path
	//fmt.Println("split fetchedURL ", fetchedURL)

	//Removing "/", prefixed to http url
	fetchedURL = strings.Replace(fetchedURL, "/", "", 1)

	//fmt.Println("replace fetchedURL ", fetchedURL)

	resp, err := http.Get(fetchedURL)
	if err != nil {
		fmt.Println("Failed to fetch the date ", err)
		return MetaData{}, err
	}
	defer resp.Body.Close()
	body, err1 := ioutil.ReadAll(resp.Body)
	if err1 != nil {
		fmt.Println("Failed to read the response body ", err)
		return MetaData{}, err
	}

	// u.Host will give the hashed part of the url
	fmt.Println("shA PART ", u.Host)

	//fetching hash of response body and verifying it by comparing it with hash provided in url( extracted via u.Host)
	valResult, err := fetchHashAndValidate(string(body), u.Host)
	if err != nil {
		fmt.Println(valResult, err)
		return MetaData{}, err
	} else {
		err1 := json.Unmarshal(body, &metadata)
		if err1 != nil {
			fmt.Println("Error while unmarshalling ", err)
			return MetaData{}, err
		}
		fmt.Println("")
		fmt.Println("Metadata")
		fmt.Println(metadata)
		return metadata, nil
	}

}

func fetchHashAndValidate(respBody string, shaPartOfURL string) (string, error) { //able to find rthe sha256 hash
	fmt.Println("fetchHashAndValidate")
	h := sha256.New()
	h.Write([]byte(respBody))
	hashedval := h.Sum(nil)

	//fmt.Printf("%x", hashedval)
	sha256Hash := "0x" + fmt.Sprintf("%x", hashedval)

	if shaPartOfURL == sha256Hash {
		return "Validation passed", nil
	} else {
		err := errors.New("Hashes did not matched")
		return "Hashes did not matched", err
	}

}
