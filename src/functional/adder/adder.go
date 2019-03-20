package main

import "fmt"

func adder() func(int) int {
	sum := 0
	return func(v int) int {
		sum += v
		return sum
	}
}

type Adder func(int) (int, Adder)

func add(base int) Adder {
	return func(i3 int) (i int, i2 Adder) {
		return 2, add(3)
	}
	return func(v int) (i int, i2 Adder) {
		return 2, add(3)
	}
}

func main() {
	a := adder()
	for i := 0; i < 10; i++ {
		fmt.Println(a(i))
	}
}
