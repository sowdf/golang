package main

import (
	"bufio"
	"fmt"
	"os"
)

func fib() func() int {
	a, b := 0, 1
	return func() int {
		a, b = b, a+b
		return a
	}
}

func writeFile(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//将一个文件写入buf 缓冲区内
	write := bufio.NewWriter(file)
	defer write.Flush()
	if err != nil {
		panic(err)
	}
	f := fib()
	for i := 0; i < 30; i++ {
		fmt.Fprintln(write, f())
	}

}

func main() {
	writeFile("abc.txt")
}
