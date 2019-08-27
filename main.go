package main

import (
	"fmt"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"log"
	"time"
)

type RedditRoot struct {
    Kind string     `json:"kind"`
    Data RedditData `json:"data"`
}

type RedditData struct {
    Children []RedditDataChild `json:"children"`
}

type RedditDataChild struct {
    Kind string `json:"kind"`
    Data *Post  `json:"data"`
}

type Post struct {
	Title		 string `json:"data"`
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

	req.Header.Set("User-Agent", "reddit-api")

	res, getErr := redditClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	type Foo struct {
		Data struct {
			Children []struct {
				Data struct {
					Title string
				}
			}
		}
	}
	
	var foo Foo
	jsonErr := json.Unmarshal(body, &foo)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	fmt.Println(foo)

}

