package spider

import (
	// "fmt"
	"github.com/qktong/go-spider/core/downloader"
)

type Spider struct {
	Name string

	Downloader downloader.Downloader

	ThreadNum int

	PagesChan chan string

	TaskChan chan string
}

func CreateSpider(name string, threadNum int) *Spider {

	spider := &Spider{Name: name, ThreadNum: threadNum, TaskChan: make(chan string, threadNum), PagesChan: make(chan string, threadNum)}

	if spider.Downloader == nil {
		spider.setDownloader(downloader.NewHttpDownloader())
	}

	return spider
}

func (this *Spider) setDownloader(d downloader.Downloader) *Spider {
	this.Downloader = d
	return this
}
func (this *Spider) getDownloader() downloader.Downloader {
	return this.Downloader
}
