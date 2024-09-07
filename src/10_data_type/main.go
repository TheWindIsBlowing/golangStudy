/**
整型、浮点型、复数、bool、字符串、byte、rune
Go语言中只有强制类型转换，没有隐式类型转换。该语法只能在两个类型之间支持相互转换的时候使用。
bool不能和整型之前互相转换，bool默认值为false

Go语言中的字符串以原生数据类型出现，使用字符串就像使用其他原生数据类型（int、bool、float32、float64 等）一样
Go 语言里的字符串的内部实现使用UTF-8编码。 字符串的值为双引号(")中的内容，可以在Go语言的源码中直接添加非ASCII码字符

Go 语言的字符有以下两种：
uint8类型，或者叫 byte 型，代表一个ASCII码字符。
rune类型，代表一个 UTF-8字符。rune类型实际是一个int32。
当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。
*/

package main

import (
	"fmt"
	"math"
)

func main() {
	// 十进制打印为二进制
	n := 10
	fmt.Printf("%b\n", n)
	fmt.Printf("%d\n", n)
	// 八进制
	m := 012
	fmt.Printf("%o\n", m)
	fmt.Printf("%d\n", m)
	// 十六进制
	q := 0x14
	fmt.Printf("%x\n", q)
	fmt.Printf("%d\n", q)

	// uint8
	var a uint8 = 255
	a += 2
	fmt.Println(a)

	// 浮点数
	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

	// 复数
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)

	// 字符串
	fmt.Println("str := \"c:\\Code\\lesson1\\go.exe\"")

	// 反引号间换行将被作为字符串中的换行，但是所有的转义字符均无效，文本将会原样输出。
	s1 := `第一行
第二行
第三行
`
	fmt.Println(s1)

}
