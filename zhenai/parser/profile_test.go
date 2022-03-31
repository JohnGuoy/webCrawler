package parser

import (
	"log"
	"testing"
	"webCrawler/fetcher"
)

func TestExtractString(t *testing.T) {
	url := "http://www.zhenai.com/zhenghun/shanghai/2"
	body, err := fetcher.Fetch(url)
	if err != nil {
		log.Printf("Fetcher: error fetching URL %s: %v", url, err)
	}
	contents := body
	/*contents, err := ioutil.ReadFile("profile_test_data.html")
	if err != nil {
		panic(err)
	}*/
	names := extractString(contents, nameRe)
	t.Log("names:", names)
	const resultSize = 20
	if len(names) != resultSize {
		t.Errorf("result length should be %d, but %d", resultSize, len(names))
	}

	genders := extractString(contents, genderRe)
	t.Log("genders:", genders)
	if len(genders) != resultSize {
		t.Errorf("result length should be %d, but %d", resultSize, len(genders))
	}

	ages := extractString(contents, ageRe)
	t.Log("ages:", ages)
	if len(ages) != resultSize {
		t.Errorf("result length should be %d, but %d", resultSize, len(ages))
	}
	/*age, err := strconv.Atoi(ages[0])
	if age != 41 {
		t.Errorf("result age should be %d, but %d", 41, age)
	}*/

	heights := extractString(contents, heightRe)
	t.Log("heights:", heights)
	if len(heights) != resultSize {
		t.Errorf("result length should be %d, but %d", resultSize, len(heights))
	}

	/*incomes := extractString(contents, incomeRe)
	t.Log("incomes:", incomes)
	if len(incomes) != resultSize {
		t.Errorf("result length should be %d, but %d", resultSize, len(incomes))
	}*/

	marriages := extractString(contents, marriageRe)
	t.Log("marriages:", marriages)
	if len(marriages) != resultSize {
		t.Errorf("result length should be %d, but %d", resultSize, len(marriages))
	}

	residences := extractString(contents, residenceRe)
	t.Log("residences:", residences)
	if len(residences) != resultSize {
		t.Errorf("result length should be %d, but %d", resultSize, len(residences))
	}

	introductions := extractString(contents, introductionRe)
	t.Log("introductions:", introductions)
	if len(introductions) != resultSize {
		t.Errorf("result length should be %d, but %d", resultSize, len(introductions))
	}

	individualInfoUrls := extractString(contents, individualInfoUrlRe)
	t.Log("individualInfoUrls:", individualInfoUrls)
	if len(individualInfoUrls) != resultSize {
		t.Errorf("result length should be %d, but %d", resultSize, len(individualInfoUrls))
	}
}

/*func TestExtractIndividualInfo(t *testing.T) {
	contents, err := ioutil.ReadFile("extractIndividualInfo_test_data.html")
	if err != nil {
		panic(err)
	}
	individualInfo1 := extractIndividualInfo(contents)
	t.Log("individualInfo1:", individualInfo1)
	if len(individualInfo1) != 7 {
		t.Errorf("result length should be %d, but %d", 8, len(individualInfo1))
	}
}

func TestParseIndividualInfoPage(t *testing.T) {
	strSlice := parseIndividualInfoPage("https://album.zhenai.com/u/1237593694")
	t.Log("strSlice:", strSlice)
	if len(strSlice) != 7 {
		t.Errorf("result length should be %d, but %d", 8, len(strSlice))
	}
}*/

func TestParseNextPage(t *testing.T) {
	url := "http://www.zhenai.com/zhenghun/shanghai/4"
	body, err := fetcher.Fetch(url)
	if err != nil {
		log.Printf("Fetcher: error fetching URL %s: %v", url, err)
	}
	contents := body

	result := parseNextPage(contents)
	t.Log("Next page URL is:", result.Url)
	nextPageUrl := "http://www.zhenai.com/zhenghun/shanghai/5"
	if result.Url != nextPageUrl {
		t.Errorf("Next page URL should be %s, but %s", nextPageUrl, result.Url)
	}
}
