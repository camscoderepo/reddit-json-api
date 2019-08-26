package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type title struct {
	Title string `json:"title"`
}


func main() {

	url := "https://www.reddit.com/r/AskReddit/.json"

	redditClient := http.Client{
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "spacecount-tutorial")

	res, getErr := redditClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	title1 := title{}
	jsonErr := json.Unmarshal(body, &title1)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(title1)

}