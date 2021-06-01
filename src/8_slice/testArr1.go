package main

import "fmt"

// fmt.Print  fmt.Println  fmt.Printf

func main() {
	// var arr [10]int // 数组长度固定 传参时要对应数组的长度
	// foreachFunc(arr)

	arr1 := [4]int{1}
	// fmt.Printf("arr 1 type = %T", arr1)
	foreachFunc(arr1)


	// 声明arr2是一个切片，并且初始化，默认值为{"liu", "feng", "hong"}，长度为3
	arr2 := []string{"liu", "feng", "hong"}
	foreachSliceFunc(arr2)

 	var slice1 []int // 声明一个切片 还没分陪空间
	slice1 = make([]int, 100) // 给切片分配空间 默认值为0
	fmt.Println(slice1)

	var slice2 []int = make([]int, 5)
	fmt.Println(slice2)
	// 声明slice3为一个切片：使用 := 推导出slice3是一个切片，为其开辟6个空间，初始化全都为0 
	slice3 := make([]int, 6)
	fmt.Println(slice3)
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
		fmt.Println(arr[i]);
	}
}