/**
数组是值类型，赋值和传参会复制整个数组。因此改变副本的值，不会改变本身的值。

数组支持 “=="、”!=" 操作符，因为内存总是被初始化过的。
[n]*T表示指针数组，*[n]T表示数组指针 。
*/

package main

import "fmt"

func main() {
	var testArray [3]int                        //数组会初始化为int类型的零值
	var numArray = [3]int{1, 2}                 //使用指定的初始值完成初始化
	var cityArray = [3]string{"北京", "上海", "深圳"} //使用指定的初始值完成初始化
	fmt.Println(testArray)                      //[0 0 0]
	fmt.Println(numArray)                       //[1 2 0]
	fmt.Println(cityArray)                      //[北京 上海 深圳]

	var testArray1 [3]int
	var numArray1 = [...]int{1, 2}
	var cityArray1 = [...]string{"北京", "上海", "深圳"}
	fmt.Println(testArray1)                          //[0 0 0]
	fmt.Println(numArray1)                           //[1 2]
	fmt.Printf("type of numArray:%T\n", numArray1)   //type of numArray:[2]int
	fmt.Println(cityArray1)                          //[北京 上海 深圳]
	fmt.Printf("type of cityArray:%T\n", cityArray1) //type of cityArray:[3]string

	// for roop
	for i := 0; i < len(cityArray1); i++ {
		fmt.Println(cityArray1[i])
	}

	// for range
	for i, v := range cityArray1 {
		fmt.Printf("%d %s\n", i, v)
	}

	// 二维数组
	cityArray2 := [3][2]string{
		{"00", "01"},
		{"10", "11"},
		{"20", "21"},
	}
	for _, v := range cityArray2 {
		for _, v := range v {
			fmt.Printf("%s\n", v)
		}
	}
}
