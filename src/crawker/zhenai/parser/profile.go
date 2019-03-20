package parser

import (
	"crawker/engine"
	"model"
	"regexp"
	"strconv"
)

var personalDetailsRe = regexp.MustCompile(`<div class="m-btn (purple|pink)" [^>]*>([^<]+)</div>`)
var ageRe = regexp.MustCompile(`"age":([\d]+)`)
var weightRe = regexp.MustCompile(`([\d]+)kg`)
var heightRe = regexp.MustCompile(`"heightString":"([\d]+)cm"`)
var hokouRe = regexp.MustCompile(`"籍贯:([^"]+)"`)
var genderRe = regexp.MustCompile(`"genderString":"([^"]+)"`)

//var nameRe  = regexp.MustCompile(`<h1 class="nickName" [^>]*>([^<]+)</h1>`)
var marriageRe = regexp.MustCompile(`"marriageString":"([^"]+)"`)
var educationRe = regexp.MustCompile(`"educationString":"([^"]+)"`)
var incomeRe = regexp.MustCompile(`"月收入:([^"]+)"`)
var carRe = regexp.MustCompile(`已买车`)
var houseRe = regexp.MustCompile("已购房")
var xingzuoRe = regexp.MustCompile(`<div class="m-btn purple" [^>]*>(..座\([^\)]+\))</div>`)

func ParseProfile(contents []byte, name string) engine.ParseResult {
	//离异 42岁 射手座(11.22-12.21) 173cm 70kg 工作地:厦门湖里区 月收入:8千-1.2万 生产/制造 大专
	profile := model.Profile{}
	profile.Name = name
	age, e := strconv.Atoi(extractString(contents, ageRe))
	if e == nil {
		profile.Age = age
	}

	height, e := strconv.Atoi(extractString(contents, heightRe))
	if e == nil {
		profile.Height = height
	}

	weight, e := strconv.Atoi(extractString(contents, weightRe))
	if e == nil {
		profile.Weight = weight
	}

	profile.Marriage = extractString(contents, marriageRe)
	//fmt.Printf("%s\n",allSubmatch)
	profile.Xingzuo = extractString(contents, xingzuoRe)
	profile.Income = extractString(contents, incomeRe)
	profile.Education = extractString(contents, educationRe)
	profile.Occupation = "不能获取"
	profile.Hokou = extractString(contents, hokouRe)
	profile.Car = extractString(contents, carRe)
	profile.Gender = extractString(contents, genderRe)
	profile.House = extractString(contents, houseRe)

	result := engine.ParseResult{Items: []interface{}{profile}}
	return result
}

func extractString(contents []byte, re *regexp.Regexp) string {
	match := re.FindSubmatch(contents)
	if len(match) >= 2 {
		return string(match[1])
	} else {
		return ""
	}
}
