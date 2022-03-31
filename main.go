package main

import (
	"webCrawler/engine"
	"webCrawler/persist"
	"webCrawler/scheduler"
	"webCrawler/zhenai/parser"
)

func main() {
	/*engine.SimpleEngine{}.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})*/

	/*e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{}, // 这个调度器运行死锁了
		WorkerCount: 10,
	}*/

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSaver(),
	}

	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
