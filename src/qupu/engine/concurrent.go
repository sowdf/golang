package engine

import (
	"fmt"
	"model"
)

type Scheduler interface {
	Submit(request Request)
	WorkerChan() chan Request
	Run()
	ReadyNotifier
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan model.Spectrum
}

func (e *ConcurrentEngine) Run(seeds ...Request) {
	out := make(chan ParseResult)

	e.Scheduler.Run()

	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}
	//提交种子
	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	for {
		result := <-out
		fmt.Printf("got item %+v", result.Item)
		/*go func() {
			//丢给外面去做 去储存什么的


			//e.ItemChan <- result.Item
		}()*/
		for _, request := range result.Requests {
			if isDuplicate(request.Url) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

var visitedUrls = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedUrls[url] {
		return true
	}
	visitedUrls[url] = false
	return false
}

/*
创建worker 去接收request
*/
func createWorker(in chan Request, out chan ParseResult, read ReadyNotifier) {
	go func() {
		for {
			read.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}
