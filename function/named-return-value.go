/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

// Function 的返回值可以被命名，被视为在 Function 顶部定义的 Variable。该 Variable 名称用于记录返回值的含义。
func sqlit(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	// 如果 return statement 不带 Parameter，则其返回值可以被命名，称为 Naked return。
	// 由于 Naked return 可能会损害 longer function 的可读性，因此仅在 short function 中使用。
	return
}

func main() {
	fmt.Println(sqlit(9))
}
