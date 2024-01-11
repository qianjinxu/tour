/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("length=%d capacity=%d %v\n", len(s), cap(s), s)
}

func main() {
	// 创建 nil Slice
	var s []int
	printSlice(s)
	if s == nil {
		fmt.Println("nil Slice\n")
	}

	// 向 nil Slice 追加 1 个 Element
	s = append(s, 0)
	printSlice(s)

	// 将单个 Element 追加到原 Slice
	s = append(s, 1)
	printSlice(s)

	// 将多个 Element 追加到原 Slice
	// 如果 s 的后备 Array 太小而无法容纳所有给定值，则会分配更大的 Array。返回的 Slice 将指向新分配的 Array。
	s = append(s, 2, 3, 4, 5, 6, 7, 8, 9)
	printSlice(s)
}
