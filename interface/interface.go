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

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

type Abser interface {
	Abs() float64
}

func main() {
	f := MyFloat(-math.Sqrt2)
	// fmt.Println(f)
	// fmt.Println(f.Abs())

	v := Vertex{3, 4}
	// fmt.Println(v)
	// fmt.Println(v.Abs())

	var a Abser
	a = f // 使用 MyFloat type 的 Variable 作为 Abser interface 的 Value
	// fmt.Println(a)
	a = v // 使用 Vertex type 的 Variable 作为 Abser interface 的 Value
	// fmt.Println(a)
	fmt.Println(a.Abs())
}
