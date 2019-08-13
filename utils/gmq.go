package utils

import (
	"sync"
	"time"
)

//go-message-queue消息队列
type Queue struct {
	size  int
	queue []Message
	sync.RWMutex
}

type Message struct {
	Id      string
	Name    string
	Time    time.Time
	Message string
}

func (this *Queue) Push(s Message) {
	this.Lock()
	defer this.Unlock()
	this.queue = append(this.queue, s)
	this.size = this.size + 1
}

func (this *Queue) PushNext(msg Message, index int) {
	this.Lock()
	defer this.Unlock()

	if index == -1 {
		this.queue = append(this.queue, msg)
	} else {
		tmp := append(this.queue[:index], msg)
		this.queue = append(tmp, this.queue[index:]...)
	}
	this.size += 1
}

//队列先进先出
func (this *Queue) Pop() Message {
	this.Lock()
	defer this.Unlock()
	s := this.queue[0]
	this.queue = this.queue[1:]
	this.size = this.size - 1
	// log.Println(s[0])
	return s
}

//队列先进后出
func (this *Queue) Pull() Message {
	this.Lock()
	defer this.Unlock()
	s := this.queue[this.size-1]
	this.queue = this.queue[0 : this.size-1]
	this.size -= 1
	return s
}

func (this *Queue) Size() int {
	this.RLock()
	defer this.RUnlock()
	return this.size
}

func (this *Queue) Get(index int) Message {
	this.RLock()
	defer this.RUnlock()
	return this.queue[index]
}

func (this *Queue) List() []Message {
	this.RLock()
	defer this.RUnlock()
	return this.queue
}

func (this *Queue) FindById(id string) (index int, msg Message) {
	this.RLock()
	defer this.RUnlock()

	index = -1
	if this.size <= 0 {
		return
	}
	for k, v := range this.queue {
		if v.Id == id {
			return k, v
		}
	}
	return
}

func (this *Queue) FindByName(name string) (index int, msg Message) {
	this.RLock()
	defer this.RUnlock()

	index = -1
	if this.size <= 0 {
		return
	}
	for k, v := range this.queue {
		if v.Name == name {
			return k, v
		}
	}
	return
}

func (this *Queue) DeleteById(id string) int {
	this.Lock()
	defer this.Unlock()

	found := -1

	for k, v := range this.queue {
		if v.Id == id {
			found = k
			break
		}
	}

	if found >= 0 {
		this.queue = append(this.queue[:found], this.queue[found+1:]...)
		this.size = this.size - 1
	}
	return found
}
