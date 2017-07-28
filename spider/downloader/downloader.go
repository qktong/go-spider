package downloader

import (
	"fmt"
)

type Downloader interface {
	Download(url string) string
}

type HttpDownloader struct {
}

func NewHttpDownloader() *HttpDownloader {
	return &HttpDownloader{}
}

func (this *HttpDownloader) Download(url string) string {
	fmt.Println("httpDownloader:", url)
	return "pages:" + url
}
