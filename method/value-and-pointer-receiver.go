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

// Abs method 有一个 v receiver，该 Receiver type 是 Vertex struct。
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Scale method 有一个 v receiver，该 Receiver type 是 *Vertex pointer。
// 使用 Pointer receiver 声明的 Method，可以修改该 Receiver 指向的 Value。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10) // 使用 Scale method 修改 *Vertex pointer 指向的 Value
	fmt.Println(v.Abs())
}
