/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

// Function 可以接受零个或多个 Parameter，每个 Parameter 由 Variable 和 Type 组成。
func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
