package main

import (
	"encoding/json"
	"log"
	"strconv"
	"ttyy/common"
	"ttyy/parser"
)

func main() {
	var dataAry []interface{}
	maxPageId := 1401
	for i := maxPageId; i > 0; i-- {
		url := "https://m.tititoy2688.com//query/book?id=" + strconv.Itoa(i)
		log.Printf("Url:%s\n", url)
		contentMap := parser.ParseList(url)
		log.Printf("%v\n", contentMap)
		log.Printf("id:%d\n", i)
		if contentMap != nil {
			log.Println("success")
			dataAry = append(dataAry, contentMap)
		}
	}

	str, err := json.Marshal(dataAry)
	if err != nil {
		panic(err)
	}
	common.CreateFile("./dataTest.json", str)
}
