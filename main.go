package main

import (
	"fmt"
	"github.com/qktong/go-spider/core/scheduler"
)

func init() {
	fmt.Println("init")

}

func main() {
	sign := make(chan int, 1)

	fmt.Println("main")
	scheduler.Start()

	sign <- 1
	<-sign
}
