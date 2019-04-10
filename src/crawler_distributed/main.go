package main

import (
	"crawler_distributed/config"
	"crawler_distributed/persist/client"
	"fmt"
	"qupu/engine"
	"qupu/qupu123/parser"
	"qupu/scheduler"
)

func main() {
	itemServer, err := client.ItemSaver(fmt.Sprintf(":%d", config.ItemSaverProt))
	if err != nil {
		panic(err)
	}
	e := engine.ConcurrentEngine{
		WorkerCount: 100,
		ItemChan:    itemServer,
		Scheduler:   &scheduler.QueuedScheduler{},
	}
	e.Run(engine.Request{
		Url:    "http://www.qupu123.com/",
		Parser: engine.NewFuncParser(parser.ParseCategoryList, "ParseCategoryList"),
	})

}
