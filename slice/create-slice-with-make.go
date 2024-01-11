/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func printSlice(s string, x []int) {
	fmt.Printf("Variable=%s, Slice length=%d, Slice capacity=%d %v\n", s, len(x), cap(x), x)
}

func main() {
	// 使用 make function 创建 Slice
	a := make([]int, 5) // 第 2 个 Parameter 表示 Slice length
	printSlice("a", a)

	b := make([]int, 0, 5) // 第 3 个 Parameter 表示 Slice capacity
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:4]
	printSlice("d", d)
}
