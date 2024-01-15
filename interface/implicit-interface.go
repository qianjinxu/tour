/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

type T struct {
	S string
}

// M() method 意味着 T type 实现了 I interface，但不需要显式声明它是这样做的。
func (t T) M() {
	fmt.Println(t.S)
}

// 隐式 Interface 是指将 Interface 的定义与实现解耦，该 Interface 可以出现在任何 Package 中而无需预先安排。
type I interface {
	M()
}

func main() {
	var i I = T{"hello"}
	i.M()
}
