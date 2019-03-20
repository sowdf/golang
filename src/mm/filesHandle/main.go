package main

import (
	"fmt"
	"net/http"
	"net/http/httputil"
	"os"
)

func PathIsExits(path string) bool {
	_, e := os.Stat(path)
	if e != nil {
		return false
	}
	if os.IsNotExist(e) {
		return false
	}
	return true
}

func getResponse(url string) interface{} {
	response, err := http.Get("http://www.4399.com")
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	dumpResponse, err := httputil.DumpResponse(response, true)
	if err != nil {
		panic(err)
	}

	return string(dumpResponse)

}

func writeFile(data string, name string) {
	path := "./" + name
	isExits := PathIsExits(path)
	if isExits {
		panic("文件已存在！！")
	}
	if !isExits {
		e := os.Mkdir(path, os.ModePerm)
		if e != nil {
			panic(e)
		}
	}
}

func main() {
	//writeFile("","czh");
	response := getResponse("http://www.4399.com")
	fmt.Println(response)

}
