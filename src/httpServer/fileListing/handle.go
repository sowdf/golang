package fileListing

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func HandleFileList(writer http.ResponseWriter, request *http.Request) error {
	path := request.URL.Path[len("/"):]
	fmt.Println(path)
	file, e := os.Open(path)
	defer file.Close()
	if e != nil {
		return e
	}
	all, e := ioutil.ReadAll(file)
	if e != nil {
		return e
	}
	writer.Write(all)
	return nil
}
