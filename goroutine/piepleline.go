package main

import (
	"fmt"
)

// 管道
// data -> natural -> squares -> dst out put..
func main() {

	//chan int 表示通道类型
	//chan<- int 单向只写通道
	//<-chan int 单向只读通道
	natural := make(chan int)
	squares := make(chan int)

	//make(chan int, 3) 容量为3的缓冲通道.

	go func() {
		for i := 0; i < 100; i++ {
			natural <- i
			//time.Sleep(1 * time.Second)
		}

		//关闭通道.
		close(natural)
	}()

	go func() {
		for {
			//ok =true 通道未关闭,正确接收数据时=true
			x, ok := <-natural
			if ok != true {
				fmt.Println("natural 通道关闭了...")
				break
			}
			squares <- x * x
		}
		close(squares)
		fmt.Println("squares 通道关闭了...")
	}()

	//range 循环上迭代通道.
	for x := range squares {
		fmt.Println(x)
	}

	// 虽然通道关闭, 依然能获取值, 只不过此时都是零值
	for {
		x, ok := <-squares
		fmt.Println(x, ok)
	}

	fmt.Println("end...")

}
