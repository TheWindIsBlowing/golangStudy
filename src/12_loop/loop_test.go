/*

Go 语言当中，所有的赋值都是拷贝：
	如果函数返回一个结构体，表示这个结构体的拷贝。
	如果函数返回一个结构体的指针，表示一个内存地址的拷贝。


*/

package loop

import (
	"fmt"
	"strings"
	"testing"
)

type person struct {
	name string
	sex  int
}

func TestStruct1(t *testing.T) {
	persons := []person{
		{
			name: "zhangsan",
			sex:  1,
		},
		{
			name: "lisi",
			sex:  2,
		},
		{
			name: "wangwu",
			sex:  1,
		},
	}
	fmt.Println(persons)

	// 值拷贝，无法修改到原切片
	for _, v := range persons {
		v.name = "_"
	}
	fmt.Println(persons)
}

func TestStruct2(t *testing.T) {
	persons := []person{
		{
			name: "zhangsan",
			sex:  1,
		},
		{
			name: "lisi",
			sex:  2,
		},
		{
			name: "wangwu",
			sex:  1,
		},
	}
	fmt.Println(persons)

	// 通过地址找到原切片进行修改
	for i := range persons {
		persons[i].name = "_"
	}
	fmt.Println(persons)
}

// range 只会在循环开始前确定一次。
func TestForRange1(t *testing.T) {
	s := []string{"hello", "world", "!"}
	fmt.Println(s)
	// 追加三次
	for range s {
		s = append(s, "_")
	}
	fmt.Println(s)
}

// 使用传统方式遍历切片，不断追加元素会导致循环永远无法结束，因为表达式 len(s) 每次循环都会重新确定一次值。
func TestForloop(t *testing.T) {
	s := []string{"hello", "world", "!"}
	fmt.Println(s)
	// 死循环
	for i := 0; i < len(s); i++ {
		s = append(s, "_")
	}
	fmt.Println(s)
}

// range中使用指针元素的影响
func TestRangePointer(t *testing.T) {
	s := Store{
		m: make(map[string]*Customer, 3),
	}
	s.storeCustomer([]Customer{
		{ID: "1", Balence: 10},
		{ID: "2", Balence: -10},
		{ID: "3", Balence: 100},
	})
	for k, v := range s.m {
		fmt.Println(k, v)
	}
}

// 跳出外层循环
func TestBreak(t *testing.T) {
myForloop:
	for i := 0; i < 5; i++ {
		fmt.Println(i)

		switch i {
		case 2:
			break myForloop
		default:
		}
	}
}

func TestString(t *testing.T) {
	s := "abc汉语123"
	// 方式一
	// range 遍历的是 s 的长度，每次遍历增加一个 code point 的长度，遍历的是每个码点的起始索引
	// 因此 s[3]无法对应字符串中的第四个完整的 code point
	for i := range s {
		fmt.Printf("position %d: %c\n", i, s[i])
	}
	fmt.Println("========================")
	// 方式二
	// 通过 range 可以直接遍历 code point（rune）
	for i, v := range s {
		fmt.Printf("position %d: %c\n", i, v)
	}
	fmt.Println("========================")
	// 方式三
	// 先将字符串转换成 rune 切片，此时遍历则会一一对应。
	// 但是转换成 rune 切片会有额外 O(N) 的空间和时间开销，如果只是希望遍历 rune，则方式二即可
	// 如果是希望获取 rune 的索引编号，则再使用方式三
	runs := []rune(s)
	for i, v := range runs {
		fmt.Printf("position %d: %c %c\n", i, v, s[i])
	}
}

// 遍历 += 的方式拼接字符串会有性能问题
func TestStringBuider(t *testing.T) {
	s := []string{"abc汉语123", "hello world", "甲乙丙"}
	total := 0
	for i := 0; i < len(s); i++ {
		total += len(s[i])
	}
	fmt.Println("s: ", s)
	fmt.Println("total: ", total)

	// 在拼接字符串之前，调用 Grow 方法，为底层字符切片分配 total 的长度空间，这样可以避免字符切片扩容造成的开销
	sb := strings.Builder{}
	sb.Grow(total)
	for _, v := range s {
		sb.WriteString(v)
	}
	fmt.Println("sb string: ", sb.String())
}
