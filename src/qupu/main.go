package main

import (
	"qupu/engine"
	"qupu/persist"
	"qupu/qupu123/parser"
	"qupu/scheduler"
)

func main() {
	itemServer, err := persist.ItemServer("qupu")
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		WorkerCount: 100,
		ItemChan:    itemServer,
		Scheduler:   &scheduler.QueuedScheduler{},
	}
	e.Run(engine.Request{
		Url:        "http://www.qupu123.com/",
		ParserFunc: parser.ParseCategoryList,
	})

}
