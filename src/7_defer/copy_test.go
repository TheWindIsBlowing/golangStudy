package deferPkg

import (
	"io"
	"os"
	"testing"
)

func TestCopy(t *testing.T) {
	CopyFile("a.txt", "b.txt")
}

func CopyFile(dstName, srcName string) (written int64, err error) {
	src, err := os.Open(srcName)
	if err != nil {
		return
	}
	defer src.Close()

	dst, err := os.Create(dstName)
	if err != nil {
		return
	}
	defer dst.Close()

	return io.Copy(dst, src)
}
