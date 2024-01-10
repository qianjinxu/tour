/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

// Constant 可以是字符、字符串、布尔值或数字，位于 Package 或 Function 级别。
// 如果 Constant 未显示指定 Type，则使用其上下文所需的 Type。
const (
	Pi = 3.14
)

func main() {
	const (
		World = "世界"
		Truth = true
	)
	fmt.Println("Hello,", World)
	fmt.Println("Happy", Pi, "Day")
	fmt.Println("Go rules?", Truth)
}
