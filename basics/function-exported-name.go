/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import (
	"fmt"
	"math"
)

func main() {
	// 如果 Function 名称以大写字母开头，则会被导出，以便其他人调用。
	// 任何未导出的 Function 名称，都无法从 Package 外部进行访问。
	fmt.Println(math.Pi)
}
