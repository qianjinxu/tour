/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	// Struct Literal 通过列出其 Field 的 Value 来表示新分配的 Struct value
	v1 = Vertex{1, 2}

	// 使用 Name: 语法仅列出 Field 的子集
	v2 = Vertex{X: 1}

	v3 = Vertex{}

	// & 操作符表示返回指向 Struct value 的 Pointer
	p = &Vertex{1, 2}
)

func main() {
	fmt.Println(v1, v2, v3, p)
}
