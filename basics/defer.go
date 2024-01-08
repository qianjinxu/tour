/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	// 延迟调用 Function，直至周围的 Function 返回后才会执行调用。
	defer fmt.Println("Qianjin Xu!")

	fmt.Printf("Hello, ")
}
