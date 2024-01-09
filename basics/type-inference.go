/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	// 如果 Variable 未显示指定 Type，则会从右侧的 Value 推断出来。
	x := 42
	y := true
	z := "hello"

	// 如果一个 Variable 显示指定 Type，则与其关联的另一个 Variable 也具有与其相同的 Type。
	a := x

	fmt.Printf("x is of type %T\n", x)
	fmt.Printf("y is of type %T\n", y)
	fmt.Printf("z is of type %T\n", z)
	fmt.Printf("a is of type %T\n", a)
}
