package downloader

import (
	"fmt"
)

type Downloader interface {
	Download(url string)
}

type HttpDownloader struct {
}

func NewHttpDownloader() *HttpDownloader {
	return &HttpDownloader{}
}

func (this *HttpDownloader) Download(url string) {
	fmt.Println("httpDownloader:", url)
}
