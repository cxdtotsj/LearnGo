package main

import (
	_ "fmt" // '_' 表示这个导入的包，可以忽略，避免编译错误
	"log"
	"time"
)

var _ = log.Println

func main() {
	_ = time.Now()
}
