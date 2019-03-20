package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "F"
	case score < 70:
		g = "D"
	case score < 80:
		g = "C"
	case score < 90:
		g = "B"
	case score < 100 || score == 100:
		g = "A"
	case score > 100 || score < 0:
		panic(fmt.Sprintf("无效分数 %d", score))
	}
	return g
}

func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s\n", contents)
	}
	fmt.Println(grade(59))
	fmt.Println(grade(69))
	fmt.Println(grade(79))
	fmt.Println(grade(89))
	fmt.Println(grade(99))
	fmt.Println(grade(100))
	//fmt.Println(grade(102));
	//fmt.Println(grade(-2));
}
