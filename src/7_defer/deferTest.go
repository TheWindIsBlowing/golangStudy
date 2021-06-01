package main

import "fmt"


// defer 关键字
// 运行顺序 压栈的形式 最后一条最先执行
// return 后的语句先执行 defer后的语句后执行
func main() {
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")

	fmt.Println("main 1")
	fmt.Println("main 2")

	testDeferAndReturn()
}

func deferFunc() {
	fmt.Println("defer func")
}

func returnFunc() int {
	fmt.Println("return func")
	return 0
}

func testDeferAndReturn() int {
	defer deferFunc()
	return returnFunc()
}

