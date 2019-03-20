package main

import (
	"fmt"
	"math"
)

var (
	aa = 3
	bb = "kk"
	cc = true
)

func variableZeroValue() {
	var s string
	var a int
	fmt.Printf("%d %q\n", a, s)
}

func variableInitialValue() {
	var a, b int = 3, 4
	var c string = "abc"
	fmt.Println(a, b, c)
}

func variableTypeDeduction() {
	var a, b, c, s = 3, 4, true, "ffff"
	fmt.Println(a, b, c, s)
}

func variableShorter() {
	a, b, c, s := 3, 4, true, "ffff"
	fmt.Println(a, b, c, s)
}

func triangle() {
	var a, b int = 3, 4
	fmt.Println(calcTriangle(a, b))
}

func calcTriangle(a, b int) int {
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	return c
}

func consts() {
	const (
		a, b     = 3, 4
		filename = "abc.txt"
	)
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(filename, c)
}

func enums() {
	const (
		cpp = iota
		python
		golang
		javascript
	)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, python, golang, javascript)
	fmt.Println(b, kb, mb, gb, tb, pb)
}
func main() {
	fmt.Println(aa)
	fmt.Println(bb)
	fmt.Println(cc)
	variableZeroValue()
	variableInitialValue()
	variableTypeDeduction()
	variableShorter()
	triangle()
	consts()
	enums()
}
