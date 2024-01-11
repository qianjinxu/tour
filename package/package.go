/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

// Go 程序由 Package 组成，而且 Go 程序必须在 main package 中开始运行。
package main

// Go 程序使用 import statement 中列出的外部 Package
import (
	"fmt"
	"math/rand"
)

func main() {
	// 按照 Go 约定，调用外部 Package 的 Function 时所使用的 Package 名称与 import path 的最后一个元素相同。
	fmt.Println("My favorite number is", rand.Intn(10))
}
