/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func do(i interface{}) {
	// 该 Switch 语句用于测试 i interface value 是否包含 int 或 string type 的 Value
	switch s := i.(type) {

	// 在每种 Type 情况下，s variable 将分别包含特定的底层 int 或 string type，并保存每种底层 Type 的 Value。
	case int:
		fmt.Printf("%v is %T type\n", s, s)
	case string:
		fmt.Printf("%v is %T type\n", s, s)
		// 在默认情况下，s variable 与 i interface value 具有相同的底层 Type 和 Value
	default:
		fmt.Printf("%v is %T type\n", s, s)
	}
}

func main() {
	do(21)
	do("hello")
	do(true)
}
