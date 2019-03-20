package main

import (
	"crawker/engine"
	"crawker/scheduler"
	"crawker/zhenai/parser"
)

/*
	乱码 安装
	gopm get -g -v golang.org/x/text
	自动获取  网页编码
	gopm get -g -v golang.org/x/net/html
*/
func main() {
	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.SimpleScheduler{},
		WorkerCount: 100,
	}
	e.Run(engine.Request{
		Url:        "http://www.zhenai.com/zhenghun",
		ParserFunc: parser.ParseCityList,
	})
}
