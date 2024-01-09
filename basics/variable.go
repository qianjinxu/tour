/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

// Variable statement 由 Variable 和 Type 组成，位于 Package 或 Function 级别。
var c, python, java bool

func main() {
	// 如果 Variable 未包含初始值，则被赋予 Zero value。不同的 Type，其 Zero value 不同。
	var i int

	fmt.Println(i, c, python, java)
}
