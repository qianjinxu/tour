/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	// defer statement 用于延迟调用 Function (e.g. fmt.Println)，直至周围的 Function (e.g. fmt.Printf) 返回后才会执行调用 Function (e.g. fmt.Println)。
	defer fmt.Println("Qianjin Xu!")

	fmt.Printf("Hello, ")
}
