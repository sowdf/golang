package main

import (
	"fmt"
	"student"
)

func main() {

	fmt.Println()
	//s1 := student.CreateStudent("张三",11);
	var s1 student.Student
	fmt.Println(s1)
	s1.SetStudentAge(11)
	s1.SetStudentName("李四")
	fmt.Println(s1)

}
