package scheduler

import (
	"webCrawler/engine"
)

type QueuedScheduler struct {
	requestChan chan engine.Request
	workerChan  chan chan engine.Request
}

func (s *QueuedScheduler) WorkerChan() chan engine.Request {
	return make(chan engine.Request)
}

func (s *QueuedScheduler) Submit(request engine.Request) {
	s.requestChan <- request
}

// WorkerReady 每个协程worker有个自己的channel，每个协程worker把“目前空闲，可以工作”消息发送给调度器
func (s *QueuedScheduler) WorkerReady(w chan engine.Request) {
	s.workerChan <- w
}

func (s *QueuedScheduler) Run() {
	s.workerChan = make(chan chan engine.Request)
	s.requestChan = make(chan engine.Request)

	go func() {
		var requestQ []engine.Request
		var workerQ []chan engine.Request
		for {
			var activeRequest engine.Request
			var activeWorker chan engine.Request
			if len(requestQ) > 0 && len(workerQ) > 0 {
				activeWorker = workerQ[0]
				activeRequest = requestQ[0]
			} /*else if len(requestQ) == 0 && len(workerQ) == 10 {
				// 程序退出条件。这里10不应该硬编码，应该读取配置文件里的Worker个数
				// 当任务来得不够快的时候，这里的条件就会满足导致程序终止运行，因此这种程序退出条件是有问题的
				os.Exit(0)
			}*/

			select {
			case r := <-s.requestChan:
				requestQ = append(requestQ, r) // request入队
			case w := <-s.workerChan:
				workerQ = append(workerQ, w) // 协程woker自己的channel入队，表示可以从里面接收消息了
			case activeWorker <- activeRequest:
				workerQ = workerQ[1:]
				requestQ = requestQ[1:]
			}
		}
	}()
}
