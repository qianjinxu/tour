/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	// 底层 Array
	names := [4]string{
		"John",
		"Paul",
		"Tom",
		"Linuz",
	}
	fmt.Println(names)

	// Slice 就像对 Array 的引用，不存储任何数据，只是描述底层 Array 的一部分。
	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)

	// 如果更改 Slice element，则会修改相应的 Array element，而且这些修改对于共享相同底层 Array 的其它 Slice 都是可见的。
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}
