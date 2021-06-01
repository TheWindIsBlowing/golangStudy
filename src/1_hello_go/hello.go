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
}

// 每个go文件都属于一个包
// golang表达式中一行结束可加;可不加  一般不加
// 函数 { 必须在同一行
// 导包建议 所有的包放括号中