package scheduler

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"github.com/qktong/go-spider/core/analyzer"
	"github.com/qktong/go-spider/core/spider"
	"github.com/qktong/go-spider/core/task"
	"github.com/qktong/go-spider/tools"
	// "strconv"
)

type SpiderConfig struct {
	Name      string
	Enable    bool
	ThreadNum int
	Index     []string
	Domain    []string
	Target    map[string]target
}

type target struct {
	Selector string
	Regexp   string
}

func init() {
	fmt.Println("sched init")
}

func Start() {
	fmt.Println("sched start")
	spidersConfig, err := getSpiderConfig()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("--------------------")
	fmt.Println(spidersConfig)

	for spiderKey, spiderConfig := range spidersConfig {
		go func(spiderKey string, spiderConfig SpiderConfig) {
			spiderName := spiderConfig.Name
			threadNum := 30
			fmt.Println("spiderKey:", spiderKey)
			fmt.Println("spiderNamexxxxxxxxxxxxxxxxxx:", spiderName)
			s := spider.CreateSpider(spiderName, threadNum)
			s.Downloader.Download("task111")
			t := task.NewTask()
			a := analyzer.NewAnalyzer()
			go t.SendTask(s)
			go t.DoTask(s)
			go a.DoAnalyze(s)
		}(spiderKey, spiderConfig)
	}
	fmt.Println("--------------------")
}

func getSpiderConfig() (spidersConfig map[string]SpiderConfig, err error) {
	spidersConfigFile, e := file.ListDir("./spiders", "toml")
	if e != nil {
		return spidersConfig, errors.New("读取爬虫配置文件失败！")
	}
	fmt.Println(spidersConfig)
	if spidersConfigFile == nil {
		fmt.Println()
		return spidersConfig, errors.New("没有爬虫的配置文件！")
	}
	var config SpiderConfig
	spidersConfig = make(map[string]SpiderConfig)
	for _, path := range spidersConfigFile {
		fileName := file.GetFileName(path)
		if _, err := toml.DecodeFile(path, &config); err != nil {
			return spidersConfig, err
		}
		spidersConfig[fileName] = config
	}
	return spidersConfig, nil
}

func AssembleSpider() {

}

func Stop(sign chan int) {
	sign <- 1
}
