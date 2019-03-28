package common

import (
	"bufio"
	"fmt"
	"os"
)

func CreateFile(filename string, content []byte) {
	//创建文件
	file, e := os.Create(filename)
	if e != nil {
		panic(e)
	}
	//关闭
	defer file.Close()

	writer := bufio.NewWriter(file)

	defer writer.Flush()

	fmt.Fprintf(writer, string(content))
}
