package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

type title struct {
	Title string `json:"title"`
}


func main() {
	resp, err := http.Get("https://www.reddit.com/r/AskReddit/")
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(body))
}