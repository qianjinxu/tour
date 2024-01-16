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

// 为 T type 声明一个 M() method 以实现 Interface，但无需将 Type 声明为实现 Interface
func (t T) M() {
	fmt.Println(t.S)
}

type I interface {
	M()
}

func main() {
	var i I = T{"hello"}
	i.M()
}
