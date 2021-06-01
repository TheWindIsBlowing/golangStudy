package main

import "fmt"

func main()  {
	fmt.Println("func foo ---------------------")
	returnA, returnB := foo(1, 2)
	fmt.Println("A, B = ", returnA, returnB)

	fmt.Println("func foo2 ---------------------")
	rA, rB := foo2(1, 2)
	fmt.Println("A, B = ", rA, rB)
}

// 多匿名返回值
func foo(a int, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return a + 1, b + 1
}

// 多有形参名称返回值
func foo2(a int, b int) (rA int, rB int)  {
	fmt.Println("a2 = ", a)
	fmt.Println("b2 = ", b)

	// rA, rB 为foo2 的形参 初始化默认为0
	fmt.Println("default rA = ", rA)
	fmt.Println("default rB = ", rB)

	rA = a * 2
	rB = b * 2

	return
}