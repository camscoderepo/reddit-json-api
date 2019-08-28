package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
	"time"
)

type List struct {
	Data struct {
		Children []struct {
			Data struct {
				Title string `json:"title"`
			}
		}
	}
}

func main() {

	
	url := "https://reddit.com/r/AskReddit.json"

	redditClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "reddit-api")

	res, getErr := redditClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	
	
	list := &List{}
	jsonErr := json.Unmarshal(body, &list)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(list)

}

