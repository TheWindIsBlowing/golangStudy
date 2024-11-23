package pkgAppend

import (
	"fmt"
	"testing"
)

type myString []string

// 为自定义的结构体或类型添加Append方法，方便使用
func (p *myString) Append(str []string) {
	slice := *p

	slice = append(slice, str...)

	*p = slice
}

func TestAppend(t *testing.T) {
	arr := myString{"hello"}
	fmt.Println(arr)
	arr.Append([]string{"world", "!"})
	fmt.Println(arr)
}

func appendTo() {

}
