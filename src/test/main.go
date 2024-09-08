package main

import (
	"fmt"
	"time"
)

var idSecond = int64(0)

func main() {
	nowSecond := time.Now().Unix()
	fmt.Println(nowSecond)
	fmt.Println(idSecond)

	a := []int{1, 3, 2, 5, 6}
	b := a[1:3]
	fmt.Println(b)

	var a1 = make([]string, 5, 10)
	for i := 0; i < 10; i++ {
		a1 = append(a1, fmt.Sprintf("%v", i))
	}
	fmt.Println(a1)
}
