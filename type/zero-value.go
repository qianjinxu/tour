/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	// 如果 Variable 未包含初始值，则被赋予 Zero value。不同的 Type，具有不同的 Zero value。
	var i int
	var b bool
	var s string
	fmt.Printf("%v %v %q\n", i, b, s)
}
