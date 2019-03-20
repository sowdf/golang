package main

import "fmt"

func main() {
	s := "内存字节"
	fmt.Println(s)

	for _, b := range []byte(s) {
		fmt.Printf("%X", b)
	}

	for i, ch := range s {
		fmt.Printf("(%d,%X)", i, ch)
	}

	for i, ch := range []rune(s) {
		fmt.Printf("(%d %c) ", i, ch)
	}
}
