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

type F float64

func (f F) M() {
	fmt.Println(f)
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

type I interface {
	M()
}

func describe(i I) {
	// 在底层，Interface Value 被认为是 Value 和具体 Type 的 Tuple。
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var i I

	i = &T{"hello"}
	describe(i)
	i.M() // 在 Interface Value 上调用 Method 会在其底层 Type 上执行同名的 Method。

	i = F(math.MaxFloat64)
	describe(i)
	i.M() // 在 Interface Value 上调用 Method 会在其底层 Type 上执行同名的 Method。
}
