package engine

import (
	"log"
	"math/rand"
	"time"
	"webCrawler/fetcher"
)

type SimpleEngine struct {
}

// Run 使用广度优先搜索URL爬取网页信息
func (e SimpleEngine) Run(seeds ...Request) {
	var requests []Request

	for _, r := range seeds {
		requests = append(requests, r)
	}

	for len(requests) > 0 {
		r := requests[0]
		requests = requests[1:]

		if r.Url == "" {
			continue
		}

		// 每来一个request就开启一个协程worker去爬
		parseResult, err := worker(r)
		if err != nil {
			continue
		}
		requests = append(requests, parseResult.Requests...)

		for _, item := range parseResult.Items {
			log.Printf("Got item %v", item)
		}

		rand.Seed(time.Now().Unix())
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)+1000))
	}

}

func worker(r Request) (ParseResult, error) {
	log.Printf("Fetching %s", r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil {
		log.Printf("Fetcher: error fetching URL %s: %v", r.Url, err)
		return ParseResult{}, err
	}

	parseResult := r.ParserFunc(body)
	return parseResult, nil
}
