/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	// 由于 C 编程语言中的 while 在 Go 编程语言中拼写为 for，因此可以省略分号。
	sum := 1
	for sum < 9 {
		sum += sum
	}

	fmt.Println(sum)
}
