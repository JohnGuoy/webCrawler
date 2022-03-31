package parser

import (
	"log"
	"regexp"
	"strconv"
	"webCrawler/engine"
	"webCrawler/model"
)

var nameRe = regexp.MustCompile(`<th><a [^>]*>([^<]+)</a></th>`)
var genderRe = regexp.MustCompile(`<td [^>]*><span [^>]*>性别：</span>([^<]+)</td>`)

// <td width="180"><span class="grayL">年龄：</span>30</td>
var ageRe = regexp.MustCompile(`<td [^>]*><span [^>]*>年龄：</span>([0-9]+)</td>`)

// <td width="180"><span class="grayL">身&nbsp;&nbsp;&nbsp;高：</span>176</td>
var heightRe = regexp.MustCompile(`<td [^>]*><span [^>]*>身[^<]*</span>([0-9]+)</td>`)

// 因为女性的月薪在城市征婚信息网页没有显示，因此男女月薪信息都到个人详细信息网页里去爬取
//var incomeRe = regexp.MustCompile(`<td><span [^>]*>月[^<]*</span>([^<]+)</td>`)
var marriageRe = regexp.MustCompile(`<td [^>]*><span [^>]*>婚况：</span>([^<]+)</td>`)
var residenceRe = regexp.MustCompile(`<td><span [^>]*>居住地：</span>([^<]+)</td>`)
var introductionRe = regexp.MustCompile(`<div class="introduce">([^<]+)</div></div>`)

var individualInfoUrlRe = regexp.MustCompile(`<a href="(http://album.zhenai.com/u/[^"]+)" target="_blank"><img [^>]*></a>`)

func ParseProfile(contents []byte) engine.ParseResult {
	/*
		https://www.zhenai.com/zhenghun/shanghai/1上海城市征婚信息第1页，征婚列表里有多个征婚人信息列表项，对于每个征婚人可以获取
		到如下个人信息：
		Name姓名（昵称）：一抹時光
		Gender性别：男士
		Residence居住地：上海
		Age年龄：43
		// Income月薪：12001-20000元
		Marriage婚况：未婚
		Height身高：163
		Introduction介绍:......
		个人详细信息网页URL：https://album.zhenai.com/u/1732017478

		个人详细信息网页里，可以获取到如下个人信息：
		Weight体重
		Education学历
		Income月薪
		Occupation职业
		Hukou籍贯
		Xinzuo星座
		House车
		Car房

		例如使用ageRe去匹配上海城市征婚信息第1页内容，匹配到的将是所有列表项里征婚人的年龄信息，因此有必要对匹配到信息进行整理。
		再进入个人详细信息网页里获取剩余个人信息。

		个人详细信息网页用XPath来匹配比较容易写，用正则表达式太难写。但是由于个人详细信息网页是需要执行js才能动态显示的，因此这部分暂时不做
		<div data-v-8b1eac0c="" class="purple-btns">
		<div data-v-8b1eac0c="" class="m-btn purple">未婚</div>
		<div data-v-8b1eac0c="" class="m-btn purple">43岁</div>
		<div data-v-8b1eac0c="" class="m-btn purple">魔羯座(12.22-01.19)</div> // Xinzuo
		<div data-v-8b1eac0c="" class="m-btn purple">163cm</div>
		<div data-v-8b1eac0c="" class="m-btn purple">58kg</div> // Weight
		<div data-v-8b1eac0c="" class="m-btn purple">工作地:上海闵行区</div>
		<div data-v-8b1eac0c="" class="m-btn purple">月收入:1.2-2万</div>
		<div data-v-8b1eac0c="" class="m-btn purple">自由职业</div> // Occupation
		<div data-v-8b1eac0c="" class="m-btn purple">高中及以下</div> // Education
		</div>
		<div data-v-8b1eac0c="" class="pink-btns">
		<div data-v-8b1eac0c="" class="m-btn pink">汉族</div>
		<div data-v-8b1eac0c="" class="m-btn pink">籍贯:内蒙古乌海</div> // Hukou
		<div data-v-8b1eac0c="" class="m-btn pink">体型:运动员型</div>
		<div data-v-8b1eac0c="" class="m-btn pink">不吸烟</div>
		<div data-v-8b1eac0c="" class="m-btn pink">稍微喝一点酒</div>
		<div data-v-8b1eac0c="" class="m-btn pink">租房</div> // House
		<div data-v-8b1eac0c="" class="m-btn pink">未买车</div> // Car
		<div data-v-8b1eac0c="" class="m-btn pink">没有孩子</div>
		<div data-v-8b1eac0c="" class="m-btn pink">是否想要孩子:想要孩子</div>
		<div data-v-8b1eac0c="" class="m-btn pink">何时结婚:时机成熟就结婚</div>
		</div>

		爬兴趣爱好和择偶条件信息的方法同爬个人资料信息，留待下个版本开发。
	*/
	names := extractString(contents, nameRe)
	genders := extractString(contents, genderRe)
	ages := extractString(contents, ageRe)
	heights := extractString(contents, heightRe)
	marriages := extractString(contents, marriageRe)
	residences := extractString(contents, residenceRe)
	introductions := extractString(contents, introductionRe)
	individualInfoUrls := extractString(contents, individualInfoUrlRe)

	length := len(names)
	if len(genders) != length || len(ages) != length || len(heights) != length || len(marriages) != length || len(residences) != length ||
		len(introductions) != length || len(individualInfoUrls) != length {
		log.Printf("ParseProfile error: inconsistent length")
		return engine.ParseResult{}
	}

	profiles := make([]model.Profile, length)

	for i := 0; i < length; i++ {
		profiles[i].Name = names[i]
		profiles[i].Gender = genders[i]

		age, err := strconv.Atoi(ages[i])
		if err == nil {
			profiles[i].Age = age
		}

		height, err := strconv.Atoi(heights[i])
		if err == nil {
			profiles[i].Height = height
		}

		profiles[i].Marriage = marriages[i]
		profiles[i].Residence = residences[i]
		profiles[i].Introduction = introductions[i]
		profiles[i].IndividualInfoUrl = individualInfoUrls[i]
		/*individualInfos := parseIndividualInfoPage(profiles[i].IndividualInfoUrl)
		if len(individualInfos) == 7 {
			profiles[i].Xinzuo = individualInfos[0]
			weight, err := strconv.Atoi(individualInfos[1])
			if err != nil {
				profiles[i].Weight = weight
			}
			profiles[i].Income=
			profiles[i].Occupation = individualInfos[2]
			profiles[i].Education = individualInfos[3]
			profiles[i].Hukou = individualInfos[4]
			profiles[i].House = individualInfos[5]
			profiles[i].Car = individualInfos[6]
		}*/
	}

	result := engine.ParseResult{}
	for _, v := range profiles {
		result.Items = append(result.Items, v)
		result.Requests = append(result.Requests, engine.Request{
			Url:        "", // 爬完个人信息后无需再爬
			ParserFunc: engine.NilParser,
		})
	}

	// 爬取“下一页”的Request
	parseNextPageResult := parseNextPage(contents)
	if parseNextPageResult.Url != "" {
		result.Requests = append(result.Requests, parseNextPageResult)
	}

	return result
}

func extractString(contents []byte, re *regexp.Regexp) []string {
	match := re.FindAllSubmatch(contents, -1)
	var strSlice []string
	for _, m := range match {
		strSlice = append(strSlice, string(m[1]))
	}

	return strSlice
}

/*func parseIndividualInfoPage(url string) []string {
	body, err := fetcher.Fetch(url)
	if err != nil {
		log.Printf("Fetcher: error fetching URL %s: %v", url, err)
		return nil
	}
	return extractIndividualInfo(body)
}

func extractIndividualInfo(contents []byte) []string {
	//ioutil.WriteFile("test.html", contents, fs.ModeAppend)
	reader := bytes.NewReader(contents)

	doc, err := htmlquery.Parse(reader)
	if err != nil {
		log.Printf("htmlquery.Parse error")
		return nil
	}

	var strSlice []string

	xinzuoNode := htmlquery.FindOne(doc, `//*[@id="app"]/div[2]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[3]`)
	strSlice = append(strSlice, htmlquery.InnerText(xinzuoNode))

	weightNode := htmlquery.FindOne(doc, `//*[@id="app"]/div[2]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[5]`)
	strSlice = append(strSlice, htmlquery.InnerText(weightNode))

	incomeNode := htmlquery.FindOne(doc, `//*[@id="app"]/div[2]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[7]`)
	strSlice = append(strSlice, htmlquery.InnerText(incomeNode))

	occupationNode := htmlquery.FindOne(doc, `//*[@id="app"]/div[2]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[8]`)
	strSlice = append(strSlice, htmlquery.InnerText(occupationNode))

	educationNode := htmlquery.FindOne(doc, `//*[@id="app"]/div[2]/div[2]/div[1]/div[2]/div/div[4]/div[1]/div[9]`)
	strSlice = append(strSlice, htmlquery.InnerText(educationNode))

	hukouNode := htmlquery.FindOne(doc, `//*[@id="app"]/div[2]/div[2]/div[1]/div[2]/div/div[4]/div[2]/div[2]`)
	strSlice = append(strSlice, htmlquery.InnerText(hukouNode))

	houseNode := htmlquery.FindOne(doc, `//*[@id="app"]/div[2]/div[2]/div[1]/div[2]/div/div[4]/div[2]/div[6]`)
	strSlice = append(strSlice, htmlquery.InnerText(houseNode))

	carNode := htmlquery.FindOne(doc, `//*[@id="app"]/div[2]/div[2]/div[1]/div[2]/div/div[4]/div[2]/div[7]`)
	strSlice = append(strSlice, htmlquery.InnerText(carNode))

	return strSlice
}*/

/*
爬取某个城市征婚信息页的下一页的URL。例如https://www.zhenai.com/zhenghun/shanghai/1上海城市征婚信息第1页的下一页URL是
http://www.zhenai.com/zhenghun/shanghai/2，直到某一页没有“下一页”元素
<a href="http://www.zhenai.com/zhenghun/shanghai/6">下一页</a>
*/
var nextPageUrlRe = regexp.MustCompile(`<a href="(http://www.zhenai.com/zhenghun/[^/]+/[0-9]+)">下一页</a>`)

func parseNextPage(contents []byte) engine.Request {
	nextPageUrl := extractNextPageUrl(contents, nextPageUrlRe)
	if nextPageUrl != "" {
		return engine.Request{
			Url:        nextPageUrl,
			ParserFunc: ParseProfile,
		}
	}
	return engine.Request{}
}

func extractNextPageUrl(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if match != nil {
		return string(match[1])
	}
	return ""
}

/*
爬取某个城市征婚信息页的下属地区的URL。例如https://www.zhenai.com/zhenghun/shanghai/1上海城市征婚信息第1页的徐汇地区URL是
https://www.zhenai.com/zhenghun/xuhui

这一块暂时不做，留待下个版本来开发。
*/
