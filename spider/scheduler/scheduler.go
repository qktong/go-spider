package scheduler

import (
	"fmt"
	"github.com/qktong/go-spider/spider/spider"
)

func init() {
	fmt.Println("sched init")
}

func Start() {
	fmt.Println("sched start")
	s := spider.CreateSpider()
	fmt.Println(s.Name)
	s.Downloader.Download("jjjj")

}

func Stop() {
	fmt.Println("sched stop")

}
