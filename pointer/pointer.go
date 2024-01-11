/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	i, j := 42, 2701

	// & 运算符表示生成指向其操作数（如 i variable）的 Pointer
	p := &i
	// *p 表示指向 p variable value 的 Pointer
	fmt.Println(*p)

	// 使用 Pointer 设置 Variable value（称为 dereferencing 或 indirecting）
	// * 运算符表示 Pointer 的 underlying value（即基础值）
	*p = 21
	fmt.Println(*p, i)

	p = &j
	*p = *p / 37
	fmt.Println(*p, j)
}
