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
	var i interface{} // i interface 是一个不包含任何 Method 签名的 empty Interface
	describe(i)

	// empty Interface 可以保存任何 Type 的 Value，但每种 Type 至少实现零个 Method。
	i = 42
	describe(i)

	i = "hello"
	describe(i)
}
