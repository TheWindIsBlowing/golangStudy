package main

import "fmt"

func main() {
	var a, b int = 10, 20
	fmt.Println("a, b = ", a, b)

	exchange(a, b)
	fmt.Println("a, b = ", a, b)


	exchange2(&a, &b)
	fmt.Println("a, b = ", a, b)

	example()
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

func example()  {
	fmt.Println("========================")
	i, j := 42, 2701

	p := &i         // 指向 i
	fmt.Println(*p) // 通过指针读取 i 的值
	*p = 21         // 通过指针设置 i 的值
	fmt.Println(i)  // 查看 i 的值

	p = &j         // 指向 j
	*p = *p / 37   // 通过指针对 j 进行除法运算
	fmt.Println(j) // 查看 j 的值
}



