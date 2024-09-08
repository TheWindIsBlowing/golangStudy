/**
要检查切片是否为空，请始终使用len(s) == 0来判断，而不应该使用s == nil来判断。
*/

package main

import "fmt"

// fmt.Print  fmt.Println  fmt.Printf

func main() {
	// var arr [10]int // 数组长度固定 传参时要对应数组的长度
	// foreachFunc(arr)

	arr1 := [4]int{1}
	// fmt.Printf("arr 1 type = %T", arr1)
	foreachFunc(arr1)

	// 声明arr2是一个切片，并且初始化，默认值为{"l", "f", "h"}，长度为3
	arr2 := []string{"l", "f", "h"}
	foreachSliceFunc(arr2)

	var slice1 []int          // 声明一个切片 还没分配空间 此时为nil
	slice1 = make([]int, 100) // 给切片分配空间 默认值为0
	fmt.Println(slice1)

	var slice2 []int = make([]int, 5)
	fmt.Println(slice2)
	// 声明slice3为一个切片：使用 := 推导出slice3是一个切片，为其开辟6个空间，初始化全都为0
	slice3 := make([]int, 6)
	fmt.Println(slice3)

	example()

	example2()
}

// 严格匹配 [4]int 类型
func foreachFunc(arr [4]int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

// 类似引用传递
func foreachSliceFunc(arr []string) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

func example() {
	fmt.Println("========================")
	s := []int{2, 3, 5, 7, 11, 13}

	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}

func example2() {
	fmt.Println("=========================")
	s := []int{2, 3, 5, 7, 11, 13, 15}
	printSlice(s)

	// 截取切片使其长度为 0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
