package main

import (
	"net/http"
	"qupu/web/controller"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir("src/qupu/web/view/static")))
	http.Handle("/api/search", controller.CreateSearchResultHandle("src/qupu/web/view/index.html"))
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		panic(err)
	}
}
