package main

import (
	"fmt"
	"retriever/mock"
	real2 "retriever/real"
	"time"
)

type Retriever interface {
	Get(url string) string
}

type Poster interface {
	Post(url string, form map[string]string) string
}

type RetrieverPost interface {
	Retriever
	Poster
}

const url = "http://www.baidu.com"

func download(r Retriever) string {
	return r.Get(url)
}

func post(poster Poster) {
	poster.Post(url, map[string]string{
		"name":   "czh",
		"course": "golang",
	})
}

func session(s RetrieverPost) string {
	s.Post(url, map[string]string{"contents": "我修改了这个mock 的值"})
	return s.Get(url)
}

func main() {
	var r Retriever
	mockRetriever := mock.Retriever{"这个时候一个mock retriever"}
	r = &mockRetriever

	inspect(r)

	fmt.Printf("%T %v\n", r, r)
	r = &real2.Retriever{UserAgent: "Mozilla/5.0", TimeOut: time.Minute}
	fmt.Printf("%T %v\n", r, r)

	if realRetriever, ok := r.(*real2.Retriever); ok {
		fmt.Println(realRetriever.UserAgent)
	} else {
		fmt.Println("not real retriever")
	}
	//fmt.Println(download(r))
	inspect(r)
	fmt.Println()
	fmt.Println(session(&mockRetriever))

}

func inspect(r Retriever) {
	fmt.Printf("%T %v", r, r)
	fmt.Println("type-switch")
	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("Contents : ", v.Contents)
	case *real2.Retriever:
		fmt.Println("UserAgent : ", v.UserAgent)
	}
}
