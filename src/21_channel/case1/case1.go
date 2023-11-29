package main

import (
	"fmt"
)

/*
channel <- value      //发送value到channel
<-channel             //接收并将其丢弃
x := <-channel        //从channel中接收数据，并赋值给x
x, ok := <-channel    //功能同上，同时检查通道是否已关闭或者是否为空
*/

func main() {
	c := make(chan int)

	go func() {
		defer fmt.Println("子go程结束")

		fmt.Println("子go程正在运行……")

		c <- 666 //666发送到c
	}()

	num := <-c //从c中接收数据，并赋值给num

	fmt.Println("num = ", num)
	fmt.Println("main go程结束")
}
