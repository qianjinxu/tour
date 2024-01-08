/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	fmt.Println("counting")

	for i := 0; i < 10; i++ {
		// 延迟调用的 Function 被压入堆栈，直到周围的 Function 返回后才会从堆栈中按照 Function 后进先出的顺序执行调用。
		defer fmt.Println(i)
	}

	fmt.Println("done.")
}
