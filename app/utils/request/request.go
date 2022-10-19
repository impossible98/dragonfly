package request

import (
	// import built-in packages
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

func Get(url string, header map[string][]string) string {
	client := http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	// control flow
	if err != nil {
		fmt.Println(err)
	}
	req.Header = header
	resp, err := client.Do(req)
	// control flow
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	// control flow
	if err != nil {
		fmt.Println(err)
	}
	// return
	return string(body)
}

func Post(url string, reqBody []byte) string {
	resp, err := http.Post(url,
		"application/json",
		bytes.NewBuffer(reqBody))
	// control flow
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	// control flow
	if err != nil {
		fmt.Println(err)
	}
	// return
	return string(body)
}
