package main

import (
	"fmt"
	"time"
)

type Demo struct {
	input         chan string
	output        chan string
	goroutine_cnt chan int
}

func NewDemo() *Demo {
	d := new(Demo)
	d.input = make(chan string, 8192)
	d.output = make(chan string, 8192)
	d.goroutine_cnt = make(chan int, 10)
	return d
}

func (this *Demo) Goroutine() {
	this.input <- time.Now().Format("2006-01-02 15:04:05")
	time.Sleep(time.Millisecond * 500)
	<-this.goroutine_cnt
}

func (this *Demo) Handle() {
	for t := range this.input {
		fmt.Println("datatime is :", t, "goroutine count is :", len(this.goroutine_cnt))
		this.output <- t + "handle"
	}
}

func main() {
	demo := NewDemo()
	go demo.Handle()
	for i := 0; i < 1000; i++ {
		demo.goroutine_cnt <- 1
		go demo.Goroutine()
	}
	close(demo.input)
}
