package engine

import (
	"math/rand"
	"time"
)

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int // 协程worker的个数
	ItemChan    chan interface{}
}

type Scheduler interface {
	Submit(request Request)
	WorkerChan() chan Request
	WorkerReady(chan Request)
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, request := range seeds {
		e.Scheduler.Submit(request)
	}

	for {
		result := <-out
		for _, item := range result.Items {
			//go func() { e.ItemChan <- item }() // 这里发生data race了，多个协程在同时读一个外部变量item，当前for循环所在的协程在
			// 写变量item，没有加锁互斥读写
			go func(item interface{}) { e.ItemChan <- item }(item) // 把外部变量item作为实参值传递给协程函数就不会造成data race了
			//e.ItemChan <- item // 也行
		}

		for _, request := range result.Requests {
			if request.Url == "" {
				continue
			}
			e.Scheduler.Submit(request)

			rand.Seed(time.Now().Unix())
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)+1000))
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}

			out <- result
		}
	}()
}
