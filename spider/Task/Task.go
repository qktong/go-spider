package task

import (
	"fmt"
	"github.com/qktong/go-spider/spider/spider"
	"math/rand"
	"strconv"
	"time"
)

type Task struct {
	ThreadNumChan chan int
	TaskChan      chan string
}

func NewTask(ThreadNum int) *Task {
	t := new(Task)
	t.TaskChan = make(chan string, ThreadNum)
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
			fmt.Println("获取任务失败， 重度次数：" + strconv.Itoa(tryTime) + " 下次重试时间： " + nextTime)
			time.Sleep(time.Second * duration)
			tryTime++
		} else {
			tryTime = 1
			fmt.Println("send a task")
			this.TaskChan <- time.Now().Format("2006-01-02 15:04:05")
		}
	}

}

func (this *Task) DoTask(s *spider.Spider) {
	for task := range this.TaskChan {
		fmt.Println(task)
		go func() {
			s.Downloader.Download(task)
		}()
	}
}
