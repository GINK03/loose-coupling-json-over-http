package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	REQ_URL = "http://localhost:8080"
)

func main() {

	client := &http.Client{Timeout: time.Duration(10 * time.Second)}
	ret := post(client, `{"type":"ADD","data":3.5}`)
	fmt.Println(ret)
}

func post(client *http.Client, value string) string {
	req, err := http.NewRequest("POST", REQ_URL+"/test", strings.NewReader(value))
	if err != nil {
		fmt.Println(err)
		return ""
	}

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	defer resp.Body.Close()

	ret := execute(resp)
	return ret
}

func execute(resp *http.Response) string {
	// response bodyを文字列で取得するサンプル
	// ioutil.ReadAllを使う
	b, err := ioutil.ReadAll(resp.Body)
	if err == nil {
		return string(b)
	}
	return ""
}

func getSimple(values url.Values) {
	// use http.Get (very simple)
	resp, err := http.Get(REQ_URL + "/get?" + values.Encode())
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	execute(resp)
}

func postSimple(values url.Values) {
	// use http.PostForm (very simple)
	resp, err := http.PostForm(REQ_URL+"/post", values)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	execute(resp)
}

func get(client *http.Client, values url.Values) {
	req, err := http.NewRequest("GET", REQ_URL+"/get", nil)
	if err != nil {
		fmt.Println(err)
		return
	}

	req.URL.RawQuery = values.Encode()

	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	execute(resp)
}
