/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	var i interface{} = "hello"

	// Assertion 用于访问 Interface value 的底层具体 Value
	s := i.(string) // 如果 Interface value 包含正确的 Type，则 Assertion 语句不会引发 Panic。
	fmt.Println(s)

	// f := i.(float64) // 如果 Interface value 未包含正确的 Type，则 Assertion 语句会引发 Panic。
	// fmt.Println(f)

	// 测试 Interface value 是否包含特定 Type
	s, ok := i.(string) // 如果 i interface value 包含正确的 string type，则 s 为 string value，ok 为 true，该 Assertion 语句不会发生 Panic。
	fmt.Println(s, ok)

	f, ok := i.(float64) // 如果 i interface value 未包含正确的 string type，则 s 为 string type 的 Zero value，ok 为 false，该 Assertion 语句不会发生 Panic。
	fmt.Println(f, ok)
}
