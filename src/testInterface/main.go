package main

import (
	"fmt"
	"testInterface/mock"
)

type Arr interface {
	Push(num int)
	Pop() int
}

func push(a Arr) {
	a.Push(1)
}

func pop(a Arr) int {
	return a.Pop()
}

func main() {
	var r Arr
	r = &mock.Arr{}
	push(r)
	fmt.Println(r)
	fmt.Println(pop(r))
	fmt.Println(r)

}
