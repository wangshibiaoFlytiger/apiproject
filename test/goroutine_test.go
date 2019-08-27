package test

import (
	"log"
	"sync"
	"testing"
	"time"
)

var wg = sync.WaitGroup{}

/**
测试协程的并发数控制: 通过sync.WaitGroup机制等待所有协程结束, 通过chan控制并发数
*/
func TestGoroutine(t *testing.T) {
	taskCount := 10
	concurrencyCount := 3
	ch := make(chan bool, concurrencyCount)
	for i := 0; i < taskCount; i++ {
		wg.Add(1)
		go execTask(ch, i)
	}

	wg.Wait()
}

/**
执行任务
*/
func execTask(ch chan bool, taskNumber int) {
	defer wg.Done()

	ch <- true
	log.Printf("执行任务: 任务号[%d], 当前时间[%d]\n", taskNumber, time.Now().Unix())
	time.Sleep(time.Second)
	<-ch
}
