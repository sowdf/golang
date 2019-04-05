package parser

import (
	"encoding/json"
	"log"
	"ttyy/fetcher"
)

func ParseList(url string) interface{} {
	contents, err := fetcher.Fetch(url)
	if err != nil {
		log.Printf("Url : %s, Error : %s", url, err)
		return ""
	}
	var obj map[string]interface{}

	err = json.Unmarshal(contents, &obj)
	if err != nil {
		log.Printf("Url : %s 解析数据失败, Error : %s", url, err)
	}
	return obj["content"]
}
