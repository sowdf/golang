package main

import (
	"crawker/engine"
	"crawker/persist"
	"crawker/scheduler"
	"crawker/zhenai/parser"
)

/*

	参考代码  ： https://github.com/jason-wj/go-zhenai-spider
	乱码 安装
	gopm get -g -v golang.org/x/text
	自动获取  网页编码
	gopm get -g -v golang.org/x/net/html
*/
func main() {
	itemChan, err := persist.ItemServer("data_profile")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 100,
		ItemChan:    itemChan,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
	/*	e.Run(engine.Request{
		Url : "http://www.zhenai.com/zhenghun/shanghai",
		ParserFunc:parser.ParseCity,
	})*/
}
