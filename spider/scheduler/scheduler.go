package scheduler

import (
	"fmt"
	"github.com/qktong/go-spider/spider/spider"
	"strconv"
	"time"
)

func init() {
	fmt.Println("sched init")
}

func Start() {
	fmt.Println("sched start")

	for crawl := 0; crawl < 1; crawl++ {
		go func(crawl int) {
			for thread_num := 0; thread_num < 3; thread_num++ {
				go func(crawl int, thread_num int) {
					s := spider.CreateSpider()
					s.Downloader.Download("===================crawl:" + strconv.Itoa(crawl) + "thread_num" + strconv.Itoa(thread_num))
				}(crawl, thread_num)
			}
		}(crawl)
		// time.Sleep(time.Second)
	}
	time.Sleep(time.Second * 10)
}

func Stop() {
	fmt.Println("sched stop")
}
