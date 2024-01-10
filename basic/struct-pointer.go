/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v

	// Struct field 可以使用 Struct pointer 访问。例如，如果要访问 Struct field X，则可以使用 (*p).X 或 p.X 表达式访问。
	p.X = 1e9

	fmt.Println((*p).X, v)
	fmt.Println(p.X, v)
}
