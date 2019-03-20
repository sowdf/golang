package main

import "fmt"

func eval(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "*":
		return a * b, nil
	case "-":
		return a - b, nil
	case "/":
		return a / b, nil
	default:
		return 0, fmt.Errorf("unsupported operation: %s", op)

	}
}

func main() {
	fmt.Println(eval(3, 4, "+"))
}
