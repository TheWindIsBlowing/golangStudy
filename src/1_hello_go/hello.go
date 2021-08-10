package main

// import "fmt"
// import "time"

import (
	"fmt"
	"time"
)

func main() {
    fmt.Println("Hello World! Welcome to Go Lang!") 
	time.Sleep(1 * time.Second);

	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	fmt.Println(today)
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}

// 每个go文件都属于一个包
// golang表达式中一行结束可加;可不加  一般不加
// 函数 { 必须在同一行
// 导包建议 所有的包放括号中