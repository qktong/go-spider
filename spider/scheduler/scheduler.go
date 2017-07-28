package scheduler

import (
	"fmt"
	"github.com/qktong/go-spider/spider/spider"
	"github.com/qktong/go-spider/spider/task"
	"time"
)

func init() {
	fmt.Println("sched init")
}

func Start() {
	fmt.Println("sched start")

	for crawl := 0; crawl < 1; crawl++ {
		go func(crawl int) {
			s := spider.CreateSpider()
			tt := task.NewTask(10)
			go tt.DoTask(s)
			go tt.SendTask(s)

		}(crawl)
		// time.Sleep(time.Second)
	}
	time.Sleep(time.Second * 10)
}

func Stop() {
	fmt.Println("sched stop")
}
