package scheduler

import (
	"fmt"
	"github.com/qktong/go-spider/spider/analyzer"
	"github.com/qktong/go-spider/spider/spider"
	"github.com/qktong/go-spider/spider/task"
	"strconv"
	"time"
)

func init() {
	fmt.Println("sched init")
}

func Start() {
	fmt.Println("sched start")
	// sign := make(chan int)
	for crawl := 0; crawl < 3; crawl++ {
		go func(crawl int) {
			spiderName := "crawl" + strconv.Itoa(crawl)
			threadNum := 30
			s := spider.CreateSpider(spiderName, threadNum)
			t := task.NewTask()
			a := analyzer.NewAnalyzer()
			go t.DoTask(s)
			go t.SendTask(s)
			go a.DoAnalyze(s)

		}(crawl)
	}
	// <-sign
}

func Stop() {
	go func() {
		time.Sleep(time.Second * 10)
		// sign <- 1
	}()
}
