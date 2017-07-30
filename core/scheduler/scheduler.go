package scheduler

import (
	"fmt"
	// "github.com/qktong/go-spider/core/analyzer"
	"github.com/qktong/go-spider/core/spider"
	// "github.com/qktong/go-spider/core/task"
	"strconv"
)

func init() {
	fmt.Println("sched init")
}

func Start() {
	fmt.Println("sched start")
	for crawl := 0; crawl < 3; crawl++ {
		go func(crawl int) {
			spiderName := "crawl" + strconv.Itoa(crawl)
			threadNum := 30
			s := spider.CreateSpider(spiderName, threadNum)
			s.Downloader.Download("task111")
			// t := task.NewTask()
			// a := analyzer.NewAnalyzer()
			// go t.SendTask(s)
			// go t.DoTask(s)
			// go a.DoAnalyze(s)
		}(crawl)
	}
}

func Stop(sign chan int) {
	sign <- 1

}
