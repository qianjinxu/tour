/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v // 等同于 sum = sum + v
	}
	c <- sum // 将 sum variable 的 Value 发送给 c channel
}

func main() {
	s := []int{
		7,
		2,
		8,
		-9,
		4,
		0,
	}

	// 声明 c channel
	c := make(chan int)

	// 对 s slice 中的数字求和：在 2 个 goroutine 之间分配工作，等待 2 个 goroutine 计算完成，则会计算最终结果。
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	x, y := <-c, <-c // 从 c channel 接收并赋值给 x, y variable

	fmt.Println(x, y, x+y)
}
