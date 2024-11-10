package deferPkg

import (
	"fmt"
	"testing"
)

func TestDefer1(t *testing.T) {
	i := 0
	j := 0
	defer func(i int) {
		fmt.Println(i, j) // j是闭包内使用外部变量，在执行到defer时确定
	}(i) // i作为匿名函数的参数传入，在一开始就确定

	i++
	j++
}

// 函数返回分为两个步骤：
// 1.为返回值赋值
// 2.return到函数调用处
// 而defer实在步骤1和步骤2之间执行
func func1() (x int) {
	x = 5
	defer func() {
		x++
	}()

	return
}

func TestDefer2(t *testing.T) {
	v := func1()
	fmt.Println("v: ", v)
}

type Person struct {
	name string
}

// 值接收者
func (p Person) print() {
	fmt.Println("person: ", p.name)
}

type Student struct {
	name string
}

// 指针接收者
func (s *Student) print() {
	fmt.Println("student: ", s.name)
}

func TestDefer3(t *testing.T) {
	p1 := Person{name: "foo"}
	s1 := Student{name: "foo"}
	defer p1.print() // 调用时已经获得一个Person的拷贝，name也已经固定
	defer s1.print() // 调用时已经获得一个Student的指针的拷贝，
	p1.name = "bar"
	s1.name = "bar"
}
