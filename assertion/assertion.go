/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// Assertion 用于提供对 Interface value 的底层 Value 的访问
	s := i.(string) // 如果 Interface value 包含特定的底层 Type，则该 Assertion 语句不会发生 Panic。
	fmt.Println(s)

	// f := i.(float64) // 如果 Interface value 未包含特定的底层 Type，则该 Assertion 语句会发生 Panic。
	// fmt.Println(f)

	// 测试 Interface value 是否包含特定的底层 Type
	s, ok := i.(string) // 如果 i interface value 包含特定的底层 string type，则 s 为 hello，ok 为 true，该 Assertion 语句不会发生 Panic。
	fmt.Println(s, ok)

	f, ok := i.(float64) // 如果 i interface value 未包含特定的底层 string type，则 s 为 0（即 string type 的 Zero value），ok 为 false，该 Assertion 语句也不会发生 Panic。
	fmt.Println(f, ok)
}
