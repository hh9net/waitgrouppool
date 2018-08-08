package waitgrouppool

import (
	"runtime"
	"sync"
)

type JobPool struct {
	wg sync.WaitGroup
	ch chan bool
}

func New(count int) *JobPool {
	c := count
	if c <= 0 {
		c = runtime.NumCPU()
	}
	return &JobPool{
		ch: make(chan bool, c),
	}
}

func (p *JobPool) Add() {
	p.ch <- true // 入队
	p.wg.Add(1)
}

func (p *JobPool) Done() {
	<-p.ch // 出队
	p.wg.Done()
}

func (p *JobPool) Wait() {
	p.wg.Wait()
}
