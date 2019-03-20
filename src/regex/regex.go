package main

import (
	"fmt"
	"regexp"
)

const text = "my email is caozhihui@4399inc.com"

func main() {
	re := regexp.MustCompile(`[0-9a-zA-z]+@[0-9a-zA-z]+\.[0-9a-zA-z]+`)
	allString := re.FindString(text)
	fmt.Println(allString)
}
