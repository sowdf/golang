package main

import (
	"net/http"
	"qupu/web/controller"
)

func main() {
	http.Handle("/", controller.CreateSearchResultHandle("src/qupu/web/view/index.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
