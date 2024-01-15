/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	// 既不包含 Value 也不包含具体 Type 的 Interface，称为 nil interface
	// 在 nil interface 上调用 Method 是一个 Go 运行时错误，因为 Interface Tuple 内没有 Type 来指示要调用哪个具体 Method。
	describe(i)
	i.M()
}
