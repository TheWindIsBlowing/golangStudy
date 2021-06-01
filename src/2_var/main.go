// 变量声明方式

package main

import (
	"fmt"
)

var name string = "fsmile"

func main() {
	// 局部变量
	var a int // 默认值为0
	fmt.Println("a = ", a)

	var b int = 100
	fmt.Println("b = ", b)

	var c = 200
	fmt.Println("c = ", c)

	d := 300 // 不支持全局变量
	fmt.Println("d = ", d)


	fmt.Println("name = ", name)


	// 单行写法
	var a1, b1, c1 int = 10, 20, 30
	fmt.Println("a1, b1, c1 = ", a1, b1, c1)
	var a2, s1 = 50, "liu"
	fmt.Println("a2, s1 = ", a2, s1)

	// 多行写法
	var (
		a3 int = 100
		s2 string = "liu1"
		bl bool = true
	)
	fmt.Println("a3, s2, bl = ", a3, s2, bl)
}