/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	m := make(map[string]int)
	fmt.Println(m)

	// 添加 Map element
	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	// 更新 Map element
	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	// 删除 Map element
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	// 测试 key 是否存在且具有 two-value 赋值
	v, ok := m["Answer"]
	// 如果 key 在 Map 中，则 ok 为 true。
	// 如果 key 在 Map 中，则 ok 为 false，且 Element 是 Map element type 的 Zero value。
	fmt.Println("The value:", v, "Present?", ok, m)
}
