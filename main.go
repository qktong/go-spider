package main

import (
	"fmt"
	"github.com/qktong/go-spider/spider/scheduler"
)

func init() {
	fmt.Println("init")

}

func main() {
	fmt.Println("main")
	scheduler.Start()

}
