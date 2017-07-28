package task

import (
	"fmt"
	"github.com/qktong/go-spider/spider/spider"
	"math/rand"
	"strconv"
	"time"
)

type Task struct {
}

func NewTask() *Task {
	t := new(Task)
	return t
}

func (this *Task) SendTask(s *spider.Spider) {
	tryTime := 1
	for true {
		num := rand.Int31n(10)
		if num < 1 {
			duration := time.Duration(1 * tryTime * tryTime)
			s, _ := time.ParseDuration("+1s")
			nextTime := time.Now().Add(s * duration).Format("2006-01-02 15:04:05")
			fmt.Println("暂时没有获取到任务 重度次数：" + strconv.Itoa(tryTime) + " 下次重试时间： " + nextTime)
			time.Sleep(time.Second * duration)
			tryTime++
		} else {
			tryTime = 1
			fmt.Println("send a task")
			s.TaskChan <- time.Now().Format("2006-01-02 15:04:05")
		}
		fmt.Println("队列名：", s.Name, "队列长度:", len(s.TaskChan))
	}

}

func (this *Task) DoTask(s *spider.Spider) {
	for task := range s.TaskChan {
		fmt.Println(task)
		time.Sleep(time.Second * 3)
		go func() {
			pageContent := s.Downloader.Download(task)
			s.PagesChan <- pageContent

		}()
	}
}
