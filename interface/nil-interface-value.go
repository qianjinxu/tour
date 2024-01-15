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

func (t *T) M() {
	if t == nil {
		fmt.Println("nil Interface value")
		return
	}
	fmt.Println(t.S)
}

type I interface {
	M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}

func main() {
	var t *T
	var i I

	// 如果 Interface 自身内部的具体 Value 为 nil，则该 Method 将使用 nil receiver 来调用。
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}
