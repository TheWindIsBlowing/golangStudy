package recoverPkg

import (
	"fmt"
	"testing"
)

func func2() {
	// recover 必须声明在 defer 函数中，因为 defer 在程序 panic 之后，依旧会执行
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recover: ", r)
		}
	}()
	func1()
	fmt.Println("ccc")
}

func func1() {
	fmt.Println("aaa")
	panic(0)
	fmt.Println("bbb")
}

func TestRecover1(t *testing.T) {
	func2()
}
