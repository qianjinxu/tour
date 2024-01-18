/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func fibo(c, quit chan int) {
	x, y := 0, 1
	for {
		// select 语句可以让 goroutine 等待多个通信操作
		// select 语句会阻塞，直至其中一个 case 可以运行，然后执行该 case。如果准备好多个 case，则会随机选择一个执行。
		select {
		case c <- x:
			x, y = y, x+y
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
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibo(c, quit)
}
