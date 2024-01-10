/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	// Slice 的 Zero value 是 nil，称为 nil Slice
	var s []int

	// nil Slice 的 Length 和 Capacity 为 0，且没有底层 Array
	fmt.Println(s, len(s), cap(s))

	// 验证 Slice 的 Zero value
	if s == nil {
		fmt.Println("This is the nil Slice.")
	}
}
