/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

// Variable 可以包含初始值，每个 Variable 各一个。
var i, j int = 1, 2

func main() {
	// 如果 Variable 包含初始值，则可以省略 Type，因为该 Variable 将采用初始值的 Type。
	var c, python, java = true, false, "no"
	fmt.Println(i, j, c, python, java)
}
