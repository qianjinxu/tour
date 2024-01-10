/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import (
	"fmt"
	"runtime"
)

func main() {
	// fmt.Println(runtime.GOOS)
	fmt.Print("Go runs on ")

	// switch statement 以 short variable 声明开始，且该 short variable 仅在 switch statement 范围内可见。
	switch os := runtime.GOOS; os {
	// 由于 Go 自动提供了每个 case 末尾所需的 break statement，因此仅运行选定的 case 而不是运行所有的 case。
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	// case 不要求是 constant，且涉及的值不要求是整数。
	default:
		fmt.Printf("%s.\n", os)
	}
}
