/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

// Function 可以返回任意数量的 Type value
func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("Hello", "Qianjin Xu!")
	fmt.Println(a, b)
}
