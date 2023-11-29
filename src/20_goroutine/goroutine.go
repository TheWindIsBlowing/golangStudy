package main

import (
	"fmt"
	"time"
)

func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new goroutine: i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}

// 主goroutine终止，其它goroutine也将终止
func main() {
	go newTask()

	i := 0
	for {
		fmt.Printf("main goroutine: i = %d\n", i)
		i++
		time.Sleep(1 * time.Second)
	}
}
