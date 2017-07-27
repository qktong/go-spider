package Task

import (
	"fmt"
)

type Task struct {
	MaxNum   int
	TaskChan chan string
}

func NewTask(maxnum int) *Task {
	t := new(Task)
	t.MaxNum = maxnum
	return t
}

func (this *Task) SendTask() {
	for true {
		if len(this.TaskChan) < this.MaxNum {
			this.TaskChan <- "x"
		}
	}
}

func (this *Task) DoTask() {
	task := <-this.TaskChan
	fmt.Println(task)
}
