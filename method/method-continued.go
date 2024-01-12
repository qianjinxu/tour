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

type Flo float64

// 在 non-Struct type 上声明 Method
// 注意：只能声明带有 Receiver argument 的 Method，该 Receiver 的 Type 与该 Receiver 必须在同一 Package 定义。 不能使用该 Receiver 的 Type 与在另一个 Package 中定义的 Receiver 来声明 Method。
func (f Flo) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := Flo(-math.Sqrt2)
	fmt.Println(f.Abs())
}
