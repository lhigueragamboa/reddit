package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
)

type Item struct {
	Title string
	URL   string
}

type Response struct {
	Data struct {
		Children []struct {
			Data Item
		}
	}
}

func main() {
	items, err := Get("golang")
	if err != nil {
		log.Fatalln(err)
	}
	for _, item := range items {
		fmt.Println(item)
	}
}

func Get(subreddit string) ([]Item, error) {
	url := fmt.Sprintf("http://reddit.com/r/%s.json", subreddit)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, errors.New(resp.Status)
	}
	r := new(Response)
	if err = json.NewDecoder(resp.Body).Decode(r); err != nil {
		return nil, err
	}
	items := make([]Item, len(r.Data.Children))
	for i, child := range r.Data.Children {
		items[i] = child.Data
	}
	return items, nil
}
