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

type Vertex struct {
	X, Y float64
}

// 将 Abs method 重写为 Abs function
func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// 将 Scale method 重写为 Scale function
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	Scale(&v, 10) // 带有 Pointer argument 的 Function，必须使用 &v 作为 Scale function 中 Vertex argument 的 Value。
	fmt.Println(Abs(v))
}
