package spider

import (
	// "fmt"
	"github.com/qktong/go-spider/spider/downloader"
)

type Spider struct {
	Name string

	Downloader downloader.Downloader

	ThreadNum uint
}

func CreateSpider() *Spider {
	// fmt.Println("create spider")

	spider := &Spider{Name: "spider name"}

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
