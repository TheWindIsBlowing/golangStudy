/*
   select {
   case <- chan1:
       // 如果chan1成功读到数据，则进行该case处理语句
   case chan2 <- 1:
       // 如果成功向chan2写入数据，则进行该case处理语句
   default:
       // 如果上面都没有成功，则进入default处理流程
   }
*/

package main

import "fmt"

func fibonacci(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
			fmt.Println("send", x)
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	fibonacci(c, quit)
}
