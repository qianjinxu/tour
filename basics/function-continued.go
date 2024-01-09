/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

// 如果两个或多个连续的 Parameter 共享同一 Type，则可以省略除最后一个 Parameter 之外的所有 Parameter 的 Type。
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
