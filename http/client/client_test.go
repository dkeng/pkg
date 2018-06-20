package client

import (
	"fmt"
	"net/url"
	"testing"
)

var (
	httpClient = New()
)

func TestGet(t *testing.T) {
	url := "https://www.baidu.com"
	body, err := httpClient.Get(url, nil)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("GET", string(body))
}

func TestPost(t *testing.T) {
	postURL := "https://www.baidu.com"
	values := url.Values{}
	values.Set("name", "keng")
	body, err := httpClient.Post(postURL, values)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("POST", string(body))
}

func TestPostJSON(t *testing.T) {
	postURL := "https://www.baidu.com"
	values := map[string]interface{}{
		"name": "keng",
	}
	body, err := httpClient.PostJSON(postURL, values)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("POST-JSON", string(body))
}
