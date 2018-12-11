package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
)

func main() {
	for _, url := range os.Args[1:] {
		filename, n, err := fetch(url)
		fmt.Println(filename)
		fmt.Println(n)
		fmt.Println(err)
	}
}

func fetch(url string) (filename string, n int64, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", 0, err
	}
	defer resp.Body.Close()
	// 获取文件名
	local := path.Base(resp.Request.URL.Path)
	fmt.Println(resp.Request.URL)
	switch local {
	case "/":
		local = "index.html"
	case ".":
		local = "index.html"
	}
	// if local == "/" or "." {
	// 	local = "index.html"
	// }
	// 创建文件
	f, err := os.Create(local)
	if err != nil {
		return "", 0, err
	}
	// 复制 body 内容至文件 f
	n, err = io.Copy(f, resp.Body)
	if closeErr := f.Close(); err == nil {
		err = closeErr
	}
	return local, n, err
}
