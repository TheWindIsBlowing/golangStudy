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
}
