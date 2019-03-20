package main

import (
	"fmt"
)

var a int
var b string
var c []float32
var d func() bool
var e struct {
	x int
}

func main() {
	var team [3]interface{}
	team[0] = 2345
	team[1] = false
	team[2] = "111"

	fmt.Print(team)
}
