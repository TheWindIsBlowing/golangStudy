package main

import (
	"fmt"
	"time"
)

/*
注意：

l channel不像文件一样需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的，才去关闭channel；

l 关闭channel后，无法向channel 再发送数据(引发 panic 错误后导致接收立即返回零值)；

l 关闭channel后，可以继续从channel接收数据；

l 对于nil channel，无论收发都会被阻塞。

可以使用 range 来迭代不断操作channel：



var ch1 chan int       // ch1是一个正常的channel，是双向的
var ch2 chan<- float64 // ch2是单向channel，只用于写float64数据
var ch3 <-chan int     // ch3是单向channel，只用于读int数据

*/

// func main() {
// 	c := make(chan int, 3) //带缓冲的通道

// 	//内置函数 len 返回未被读取的缓冲元素数量， cap 返回缓冲区大小
// 	fmt.Printf("len(c)=%d, cap(c)=%d\n", len(c), cap(c))

// 	go func() {
// 		defer fmt.Println("子go程结束")

// 		for i := 0; i < 3; i++ {
// 			c <- i
// 			fmt.Printf("子go程正在运行[%d]: len(c)=%d, cap(c)=%d\n", i, len(c), cap(c))
// 		}
// 	}()

// 	time.Sleep(2 * time.Second) //延时2s
// 	for i := 0; i < 3; i++ {
// 		num := <-c //从c中接收数据，并赋值给num
// 		fmt.Println("num = ", num)
// 	}
// 	fmt.Println("main进程结束")
// }

// func main() {
// 	c := make(chan int)

// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			c <- i
// 		}

// 		time.Sleep(2 * time.Second)
// 		close(c)
// 	}()

// 	for {
// 		//ok为true说明channel没有关闭，为false说明管道已经关闭
// 		if data, ok := <-c; ok {
// 			fmt.Println(data)
// 		} else {
// 			break
// 		}
// 	}

// 	fmt.Println("Finished")
// }

func main() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		time.Sleep(2 * time.Second)
		close(c)
	}()

	for data := range c {
		fmt.Println(data)
	}
	fmt.Println("Finished")
}
