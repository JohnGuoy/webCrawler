package parser

import (
	"io/ioutil"
	"testing"
)

// 测试ParseCityList函数的功能是否符合预期
func TestParseCityList(t *testing.T) {
	contents, err := ioutil.ReadFile("citylist_test_data.html")
	if err != nil {
		panic(err)
	}

	result := ParseCityList(contents)

	const resultSize = 24
	if len(result.Requests) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Requests))
	}
	if len(result.Items) != resultSize {
		t.Errorf("result should have %d requests, but had %d", resultSize, len(result.Items))
	}

	expectedUrls := []string{
		"http://www.zhenai.com/zhenghun/beijing",
		"http://www.zhenai.com/zhenghun/shanghai",
		"http://www.zhenai.com/zhenghun/guangzhou",
	}
	for i, url := range expectedUrls {
		if result.Requests[i].Url != url {
			t.Errorf("expected url #%d: %s, but was %s", i, url, result.Requests[i].Url)
		}
	}

	expectedCities := []string{
		"北京", "上海", "广州",
	}
	for i, city := range expectedCities {
		if result.Items[i].(string) != city {
			t.Errorf("expected url #%d: %s, but was %s", i, city, result.Items[i])
		}
	}
}
