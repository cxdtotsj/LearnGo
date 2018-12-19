// package main

// import (
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// 	"net/url"
// )

// const baseURL string = "https://dt-dev.arctron.cn/api"

// // const ContentType = "application/x-www-form-urlencoded"

// type body struct {
// 	email    string
// 	password string
// }

// func main() {
// 	api := "/user/login"
// 	// data := "email=demo,password=123456"
// 	// data := `{
// 	// 	"eamil":"demo",
// 	// 	"password":"123456"
// 	// }`
// 	data := url.Values{"email": {"demo"}, "password": {"123456"}}
// 	body := hTTPPost(api, data)
// 	fmt.Println(body)
// }

// func hTTPPost(api string, data url.Values) string {

// 	url := baseURL + api
// 	// params := strings.NewReader(data)
// 	// resp, err := http.Post(url, ContentType, params)
// 	resp, err := http.PostForm(url, data)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	body, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	return string(body)
// }
