package analyzer

import (
	"fmt"
	"github.com/qktong/go-spider/core/spider"
)

type Analyzer struct {
}

func NewAnalyzer() *Analyzer {
	a := new(Analyzer)
	return a
}

func (this *Analyzer) DoAnalyze(s *spider.Spider) {
	for page := range s.PagesChan {
		go func() {
			fmt.Println(page)
		}()
	}
}
