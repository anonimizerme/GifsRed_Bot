package redgifs

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Response struct {
	Gifs []Gif `json:"gifs"`
}

type Gif struct {
	Id   string   `json:"id"`
	Tags []string `json:"tags"`
	Urls Urls     `json:"urls"`
}

type Urls struct {
	HDUrl      string `json:"hd"`
	Thumbnail  string `json:"thumbnail"`
	VThumbnail string `json:"vthumbnail"`
}

func Search(query string) (*Response, error) {
	res, err := http.Get("https://api.redgifs.com/v2/gifs/search?type=g&search_text=" + query + "&order=trending&count=10")

	if err != nil {
		fmt.Printf("error %v", err)
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		fmt.Printf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, err
	}

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("error %v", err)
		return nil, err
	}

	var body Response

	err = json.Unmarshal(resBody, &body)
	if err != nil {
		fmt.Printf("error %v", err)
		return nil, err
	}

	return &body, nil
}
