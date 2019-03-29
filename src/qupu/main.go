package main

import (
	"qupu/engine"
	"qupu/qupu123/parser"
	"qupu/scheduler"
)

func main() {
	e := engine.ConcurrentEngine{
		WorkerCount: 100,
		Scheduler:   &scheduler.QueuedScheduler{},
	}
	e.Run(engine.Request{
		Url:        "http://www.qupu123.com/",
		ParserFunc: parser.ParseCategoryList,
	})

}
