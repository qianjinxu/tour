/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(2)        // 调用具有 Pointer receiver 的 Method 时，采用 Value 或 Pointer 作为 Receiver
	ScaleFunc(&v, 10) // 调用带有 Pointer argument 的 Function 时，必须采用 Pointer

	p := &Vertex{4, 3}
	p.Scale(3)      // 调用具有 Pointer receiver 的 Method 时，采用 Value 或 Pointer 作为 Receiver
	ScaleFunc(p, 8) // 调用带有 Pointer argument 的 Function 时，必须采用 Pointer

	// 对于 v.Scale(2) 和 p.Scale(3) statement 来讲，即使 v 是一个 Value 而不是一个 Pointer，也会自动调用具有 Pointer receiver 的 Method。
	// 也就是说，Go 会将 v.Scale(2) statement 解释为 (&v).Scale(2) statement，因为 Scale method 有一个 Pointer receiver。
	fmt.Println(v, p)
}
