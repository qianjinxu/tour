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

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func AbsFunc(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())    // 调用具有 Value receiver 的 Method 时，采用 Value 或 Pointer 作为 Receiver。在这种情况下，调用 v.Abs() method 被解释为 (*v).Abs()
	fmt.Println(AbsFunc(v)) // 调用带有 Value argument 的 Function 时，必须采用该特定 Type 的 Value
}
