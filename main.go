package main

import (
	"fmt"
	"github.com/qktong/go-spider/spider/scheduler"
)

func init() {
	fmt.Println("init")

}

func main() {
	sign := make(chan int, 1)

	fmt.Println("main")
	scheduler.Start()

	<-sign
}
