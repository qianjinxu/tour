/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// 在 Function 内部，可以使用 short assignment statement（即 := 符号）来代替隐式的 var statement。
	// 在 Function 外部，不可以使用 short assignment statement。
	k := 3
	c, python, java := true, false, "no"

	fmt.Println(i, j, k, c, python, java)
}
