package main

import "fmt"

func main() {

	m := map[string]string{
		"name": "czh",
		"cccc": "cccc",
	}

	fmt.Println(m)

	m2 := make(map[string]int)

	fmt.Println(m2)

	var m3 map[string]int
	fmt.Println(m3)

	for k, v := range m {
		fmt.Println(k, v)
	}

	name := m["name"]
	fmt.Println(name)

	if test, ok := m["cccc"]; ok {
		fmt.Println(test)
	} else {
		fmt.Println("key no exits")
	}

	delete(m, "cccc")

	fmt.Println(m)

}
