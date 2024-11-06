package optionargs

import (
	"fmt"
	"testing"
)

func NewServer(addr string, opts ...Option) (string, error) {
	// 有错误直接返回
	var options options
	for _, opt := range opts {
		err := opt(&options)
		if err != nil {
			return "", err
		}
	}

	// 字段配置
	port := options.port
	content := options.content
	fmt.Println(port, content)

	return "success", nil
}

func TestRun(t *testing.T) {
	str, err := NewServer("", WithPort(456), WithContent([]string{"gggg", "ahaha"}))
	fmt.Println(str)
	fmt.Println(err)
}
