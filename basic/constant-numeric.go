/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

// 如果 Constant 未显示指定 Type，则使用其上下文所需的 Type。
const (
	// 左移 100 位（即 1 后面跟 100 个 0 的二进制数字）
	Big = 1 << 100
	// 右移 99 位（即 1 后面跟 1 个 0 的二进制数字）
	Small = Big >> 99
)

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needFloat(Big))
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
}
