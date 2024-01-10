/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("sliceLength=%d, sliceCapacity=%d, sliceElement=%v\n", len(s), cap(s), s)
}

// 示例说明：每次只取消一组注释，来观察 Slice length 和 Slice capacity 的变化
func main() {
	s := []int{1, 2, 3, 4, 5}

	// Slice length 是 5, Slice capacity 是 5
	// s = s[:]
	// printSlice(s)

	// Slice length 是 0, Slice capacity 是 5
	// s = s[:0]
	// printSlice(s)

	// Slice length 是 4, Slice capacity 是 5
	// s = s[:4]
	// printSlice(s)

	// Slice length 是 2, Slice capacity 是 2
	// s = s[3:]
	// printSlice(s)

	// Slice length 是 1, Slice capacity 是 2
	// s = s[3:4]
	// printSlice(s)
}
