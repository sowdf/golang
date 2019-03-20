package main

import "fmt"

type Student struct {
	name string
	age  int
	sid  int
	arr  []int
}

func main() {
	var student Student
	student.name = "张三"
	student.age = 18
	student.sid = 9527
	student.arr = append(student.arr, 11333)
	fmt.Println(student)
}
