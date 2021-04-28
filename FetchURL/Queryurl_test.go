package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	s "strings"
	"testing"
)

var metaDataUrl = flag.String("metaDataUrl", "http://www.google.com", "url")

/*
go test -run TestQueryHttpUrl -metaDataUrl "https://storage.googleapis.com/tzip-16/emoji-in-metadata.json"

go test -run TestQuerySHA256Url -metaDataUrl "sha256://0x7e99ecf3a4490e3044ccdf319898d77380a2fc20aae36b6e40327d678399d17b/https:%2F%2Fstorage.googleapis.com%2Ftzip-16%2Ftaco-shop-metadata.json"

*/

func TestQuerySHA256Url(t *testing.T) {

	fmt.Println("testng TestSHA256")

	metaDataURL := *metaDataUrl

	fmt.Println(metaDataURL)

	metaDataURL = s.Split(s.Split(metaDataURL, "://")[1], "/")[1]

	if s.Contains(metaDataURL, "%2F%2F") {
		metaDataURL = s.Replace(metaDataURL, "%2F%2F", "//", -1)
		if s.Contains(metaDataURL, "%2F") {
			metaDataURL = s.Replace(metaDataURL, "%2F", "/", -1)
		}
	}

	resp, err := http.Get(metaDataURL)
	if err != nil || resp.StatusCode != 200 {
		t.Errorf("Error while querying the url.Error is  %d. Response code %d\n", err, resp.StatusCode)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Error while reading the response body.Error is  %d.", err)
		}
		var metadata MetaData
		err1 := json.Unmarshal(body, &metadata)
		if err1 != nil {
			t.Errorf("Error while unmarshalling the response.Error is  %d.", err)
		} else {
			fmt.Println("MetaData fetched ", metadata)
		}

	}

}

func TestQueryHttpUrl(t *testing.T) {

	fmt.Println("testng TestSHA256")

	metaDataURL := *metaDataUrl

	fmt.Println(metaDataURL)

	resp, err := http.Get(metaDataURL)
	if err != nil || resp.StatusCode != 200 {
		t.Errorf("Error while querying the url.Error is  %d. Response code %d\n", err, resp.StatusCode)
	} else {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			t.Errorf("Error while reading the response body.Error is  %d.", err)
		}
		var metadata MetaData
		err1 := json.Unmarshal(body, &metadata)
		if err1 != nil {
			t.Errorf("Error while unmarshalling the response.Error is  %d.", err1)
		} else {
			fmt.Println("MetaData fetched ", metadata)
		}

	}

}
