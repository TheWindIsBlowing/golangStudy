package syncPkg

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var i int
var mutex = sync.Mutex{}

// i++的操作可以被分解成三个步骤
// 1.读取i的值
// 2.对应值++
// 3.将值写回i
// 当并发执行两个协程的时候，i 的最终结果是无法预计的
func add(i *int) {
	go func() {
		for a := 0; a < 5; a++ {
			*i++
		}
	}()
}
func TestFunc(t *testing.T) {
	i := 0

	go func() {
		// for a := 0; a < 5; a++ {
		i++
		// }
	}()
	go func() {
		// for a := 0; a < 5; a++ {
		i++
		// }
	}()

	time.Sleep(time.Second)
	fmt.Println("i: ", i)
}

func add1(i *int) {
	go func() {
		mutex.Lock()
		*i++
		mutex.Unlock()
	}()
}
func TestFunc1(t *testing.T) {
	i := 0

	add1(&i)
	add1(&i)

	time.Sleep(time.Second)
	fmt.Println("i: ", i)
}

func TestFunc2(t *testing.T) {
	i := 0
	ch := make(chan struct{})
	go func() {
		fmt.Println("aaa")
		i = 1
		<-ch
		fmt.Println("bbb")
	}()
	time.Sleep(time.Second)
	ch <- struct{}{}
	time.Sleep(time.Second)
	fmt.Println("i: ", i)
}
