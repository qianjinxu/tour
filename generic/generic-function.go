/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

// Mygeneric function 是一个 Generic function
func Mygeneric[T comparable](s []T, x T) int {
	for i, v := range s {
		// 由于 v 和 x parameter 均为 T type，因此具有 T type 的 comparable type constraint。
		// 使用 comparable type constraint 将 x value 与 s slice element 进行逐个比较
		if v == x {
			// 如果找到匹配项，则返回 s slice element 中 x value 的 Index
			return i
		}
	}
	// 如果未找到匹配项，则返回 -1。
	return -1
}

func main() {
	// Index 适用于 int type 的 Slice
	si := []int{
		10,
		20,
		15,
		-10,
	}
	fmt.Println(Mygeneric(si, 20))

	// Index 也适用于 string type 的 Slice
	ss := []string{
		"foo",
		"bar",
		"baz",
	}
	fmt.Println(Mygeneric(ss, "hello"))
}
