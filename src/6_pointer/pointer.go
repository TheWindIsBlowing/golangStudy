package main

import "fmt"

func main() {
	var a, b int = 10, 20
	fmt.Println("a, b = ", a, b)

	exchange(a, b)
	fmt.Println("a, b = ", a, b)


	exchange2(&a, &b)
	fmt.Println("a, b = ", a, b)
}

// 错误写法 交换两个数的值
func exchange(a, b int) {
	temp := a;
	a = b;
	b = temp;
}

// 指针写法
func exchange2(a, b *int) {
	temp := *a
	*a = *b
	*b = temp
}



