/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

// type I interface {
// }

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	// 当处理未知 Type 的 Value 时，则可以使用 empty Interface。
	var i interface{} // interface{} 表示一个 empty Interface（即不包含任何 Method 的 Interface）
	describe(i)

	// empty Interface 可以保存任何 Type 的 Value，但每种 Type 至少实现零个 Method。
	i = 42
	describe(i)

	i = "hello"
	describe(i)
}
