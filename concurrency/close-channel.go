/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import (
	"fmt"
	"time"
)

func fibo(t int, c chan int) {
	x, y := 0, 1
	for i := 0; i < t; i++ {
		c <- x
		fmt.Printf("x=%v, y=%v\n", x, y)
		x, y = y, x+y
	}
	close(c) // 关闭 c channel 以指示不再发送任何 Value
}

func main() {
	c := make(chan int, 10)
	// fmt.Printf("%v\n", c)
	go fibo(cap(c), c)
	// fmt.Printf("%d\n", cap(c))

	// 重复从 c channel 接受 Value，直至该 Channel 关闭。
	for v := range c {
		fmt.Println(v, time.Now())
	}
}
