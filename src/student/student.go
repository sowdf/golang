package student

import "fmt"

type Student struct {
	Name string
	Age  int
}

func (student Student) PointStudent() {
	fmt.Println(student.Name)
}

func (student *Student) SetStudentName(Name string) {
	student.Name = Name
}

func (student *Student) SetStudentAge(Age int) {
	student.Age = Age
}

func CreateStudent(Name string, Age int) *Student {
	return &Student{
		Name: Name,
		Age:  Age,
	}
}
