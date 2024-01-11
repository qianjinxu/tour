/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4

	// Type 转换通过使用 T(v) 表达式将 v value 转换为 T type
	// 与 C 不同，在 Go 中不同 Type 的 item 之间的赋值需要显示转换。
	var f float64 = math.Sqrt(float64(x*x + y*y))
	var z uint = uint(f)

	fmt.Println(x, y, z)
}
