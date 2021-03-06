package parser

import (
	"regexp"
	"webCrawler/engine"
)

const cityListRe = `<a target="_blank" href="(http://www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

//const cityListRe = `<a target="_blank" href="//(www.zhenai.com/zhenghun/[0-9a-z]+)" [^>]*>([^<]+)</a>`

func ParseCityList(contents []byte) engine.ParseResult {
	re := regexp.MustCompile(cityListRe)
	matches := re.FindAllSubmatch(contents, -1)

	result := engine.ParseResult{}
	for _, m := range matches {
		result.Items = append(result.Items, string(m[2])) // m[2]是城市名字
		result.Requests = append(result.Requests, engine.Request{
			Url:        string(m[1]), // 城市名字对应的URL
			ParserFunc: ParseProfile, // 爬该城市主页下征婚人信息的回调函数
		})
		// fmt.Printf("City: %s, URL: %s\n", m[2], m[1])
	}
	// fmt.Printf("Matches found: %d\n", len(matches))

	return result
}
