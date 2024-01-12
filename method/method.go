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

// Method 是带有 Receiver argument 的特殊 Function
// Abs() method 有一个 v receiver，该 Receiver 是 Vertex type。
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v)
	fmt.Println(v.Abs())
}
