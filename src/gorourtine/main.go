package main

import (
	"fmt"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				//没有人交出控制权
				fmt.Printf("hello form goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)
}
