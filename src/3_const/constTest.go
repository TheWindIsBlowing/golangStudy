
package main

import "fmt"

// const 来表示枚举
const(
	// 关键字 iota 第一行默认值为0 每行递增1
	BEIJING = iota // iota 微量;极少量;希腊字母表的第9个字母
	SHANGHAI // iota = 1
	GUANGZHOU // iota = 2
	SHENZHEN
)

const(
	a, b = iota * 1, iota * 2
	c, d
	e, f

	g, h = iota + 3, iota + 2
	i, j
	k, l
)

func main() {
    fmt.Println("const study")

	const aa int = 100
	fmt.Println("const aa = ", aa)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHAHAI = ", SHANGHAI)
	fmt.Println("GUANZHOU = ", GUANGZHOU)
	fmt.Println("SHENZHEN = ", SHENZHEN)

	fmt.Println("a, b = ", a, b)
	fmt.Println("c, d = ", c, d)
	fmt.Println("e, f = ", e, f)
	fmt.Println("g, h = ", g, h)
	fmt.Println("i, j = ", i, j)
	fmt.Println("k, l = ", k, l)
}