package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {

	// make constant since you're not changing this value.
	// also, it should be AskReddit.json
	// also, www's aren't really a thing anymore since it's technically
	// a subdomain so it'll usually be re-routed to just https://reddit.com
	url := "https://www.reddit.com/r/AskReddit/.json"

	// it's really good you caught the timeout stuff
	redditClient := http.Client{
		// probably want timeout to be a constant as well but that's personal
		// preference
		Timeout: time.Second * 2, // Maximum of 2 secs
	}

	// nice
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	// gj knowing to set the UA
	req.Header.Set("User-Agent", "reddit-api")

	res, getErr := redditClient.Do(req)
	if getErr != nil {
		log.Fatal(getErr)
	}
	// this is a little tricky and something that you just need to remember
	// but when doing an http request it creates a pipe and you need to make
	// sure that the pipe closes when this is all done otherwise it'll create
	// a memory leak (at least if this is a long-running service).  Defer is
	// also a special keyword in that basically it says "at the end of this
	// function, no matter what the return is, it'll call the defered stuff
	// before the function actually returns".  Basically, even if there is a
	// log.Fatal(), it'll call the defered stuff first.
	defer res.Body.Close()

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}
	// looks good but in-line structs (meaning defined in the function)
	// are generally not the way to do it because then you limit the scope
	// to this function.  So put this on something above the main func.

	// Also, I haven't ran your code but it looks like this should work.
	// however, don't rely on the json name to be equal to your field
	// name.  Always add the json tag eg...
	type Foo struct {
		Data struct {
			Children []struct {
				Data struct {
					Title string // `json="title"`, or whatever the json field name is
				}
			}
		}
	}

	// do foo := &foo{}.  This way you set the variable to be a pointer to
	// instead value of (generally you want all structs to be pointers
	// so you don't have to pass the value around in memory)
	var foo Foo
	jsonErr := json.Unmarshal(body, &foo)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}

	// I haven't ran the code but I think you want something like:
	/*
		for _, post := range foo {
			if count(post.Data.Children) == 0 {
				log.Info("post doesn't contain children")
				continue
			}
			fmt.Println(post.Data.Children[0].Data.Title)
		}
	*/
	// the reason is that you'll have more info than just the Title

	fmt.Println(foo)

}
