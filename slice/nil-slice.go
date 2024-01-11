/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	// non-nil Slice
	var nonNilSlice = []int{
		1,
		2,
		3,
	}
	fmt.Println(len(nonNilSlice), cap(nonNilSlice), nonNilSlice)
	if nonNilSlice != nil {
		fmt.Println("non-nil Slice \n")
	}

	// nil Slice
	var nilSlice []int                                  // Slice 的 Zero value 是 nil（称为 nil Slice）
	fmt.Println(len(nilSlice), cap(nilSlice), nilSlice) // nil Slice 的 Length 和 Capacity 均为 0，且没有底层 Array。
	if nilSlice == nil {                                // 验证 Slice 的 Zero value
		fmt.Println("nil Slice \n")
	}
}
