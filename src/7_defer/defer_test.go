package deferPkg

import (
	"fmt"
	"testing"
)

// defer 关键字
// 运行顺序 压栈的形式 最后一条最先执行
// return 后的语句先执行 defer后的语句后执行
func TestDefer(t *testing.T) {
	// defer fmt.Println("defer 1")
	// defer fmt.Println("defer 2")

	// fmt.Println("main 1")
	// fmt.Println("main 2")

	// testDeferAndReturn()

	a()
	b()
	fmt.Println(c())
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

// The behavior of defer statements is straightforward and predictable. There are three simple rules:
// 1. A deferred function's arguments are evaluated when the defer statement is evaluated.
func a() {
	i := 0
	defer fmt.Println(i)
	i++
	return
}

// This function prints "0"
// 2. Deferred function calls are executed in Last In First Out order after the surrounding function returns.
func b() {
	defer fmt.Println()
	for i := 0; i < 4; i++ {
		defer fmt.Print(i)
	}
}

// This function prints "3210"
// 3. Deferred functions may read and assign to the returning function's named return values.
func c() (i int) {
	defer func() { i++ }()
	return 1
}

// This function returns 2
