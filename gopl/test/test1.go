package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const baseUrl = "https://dt-dev.arctron.cn/api"
const ContentType = "x-www-form-urlencoded"

type body struct {
	email    string
	password string
}

func main() {
	api := "/user/login"
	url := baseUrl + api
	data := strings.NewReader("email=demo,password=123456")
	resp, err := http.Post(url, ContentType, data)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", body)
}

func HttpPost(api, data string) string {

	url := baseUrl + api
	params := strings.NewReader(data)
	resp, err := http.Post(url, ContentType, params)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp)
	if err != nil {
		log.Fatal(err)
	}

	return string(body)
}
