package engine

type Request struct {
	Url        string                   // URL
	ParserFunc func([]byte) ParseResult // 爬该URL的回调函数
}

type ParseResult struct {
	Requests []Request
	Items    []interface{} // 存储爬到的信息条目的
}

// NilParser 默认使用的ParserFunc函数
func NilParser([]byte) ParseResult {
	return ParseResult{}
}
