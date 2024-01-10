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
		// 将所有延迟调用 Function 放入堆栈，直到周围的 Function 返回后，才会从堆栈中按照其后进先出的顺序执行调用。
		defer fmt.Println(i)
	}

	fmt.Println("done.")
}
