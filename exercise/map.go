/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

// 任务说明：https://go.dev/tour/moretypes/23
package main

import (
	"golang.org/x/tour/wc"
)

// 统计字数
func WordCount(s string) map[string]int {
	return map[string]int{"x": 1}
}

func main() {
	// 应返回 s string 中每个 word 计数的 Map
	wc.Test(WordCount)
}
