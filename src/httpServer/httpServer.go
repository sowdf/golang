package main

import (
	"httpServer/fileListing"
	"log"
	"net/http"
	"os"
)

type appHandle func(writer http.ResponseWriter, request *http.Request) error

func errorWrapper(app appHandle) func(writer http.ResponseWriter, request *http.Request) {
	return func(writer http.ResponseWriter, request *http.Request) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("Panic : %v", r)
				http.Error(writer, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
		}()
		e := app(writer, request)
		if e != nil {
			code := http.StatusOK
			switch {
			case os.IsNotExist(e):
				code = http.StatusNotFound
			case os.IsPermission(e):
				code = http.StatusForbidden
			default:
				code = http.StatusInternalServerError
			}
			http.Error(writer, http.StatusText(code), code)
		}
	}
}

func main() {
	http.HandleFunc("/", errorWrapper(fileListing.HandleFileList))
	err := http.ListenAndServe(":8889", nil)
	if err != nil {
		panic(err)
	}
}
