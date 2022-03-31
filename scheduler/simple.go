package scheduler

import "webCrawler/engine"

type SimpleScheduler struct {
	workerChan chan engine.Request
}

func (s *SimpleScheduler) WorkerChan() chan engine.Request {
	return s.workerChan
}

func (s *SimpleScheduler) WorkerReady(chan engine.Request) {

}

func (s *SimpleScheduler) Run() {

}

func (s *SimpleScheduler) Submit(request engine.Request) {
	// 把爬取请求发送给Channel。这里若不开一个协程来发送，会造成循环等待死锁
	go func() { s.workerChan <- request }()
}
