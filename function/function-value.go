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

func compute(fn func(float64, float64) float64) float64 {
	return fn(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y) // math.Sqrt function 用于返回 x 的平方根
	}

	// 求平方根
	fmt.Println(hypot(5, 12)) // 169=13²

	// 将 compute function 用作 hypot function 的 Parameter
	fmt.Println(compute(hypot)) // 25=5²

	// 将 compute function 用作 math.Pow function 的 return value
	fmt.Println(compute(math.Pow)) // math.Pow function 用于返回 x**y（如 3⁴=81）
}
