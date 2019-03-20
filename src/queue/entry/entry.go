package main

import (
	"fmt"
	"queue"
)

func main() {
	a := queue.Queue{1}
	a.Push(2)
	a.Push(3)
	//a.Push("absfsfs ")
	fmt.Println(a.Pop())
	fmt.Println(a.Pop())
	fmt.Println(a)
	fmt.Println(a.IsEmpty())
	fmt.Println(a.Pop())
	fmt.Println(a.IsEmpty())
}
