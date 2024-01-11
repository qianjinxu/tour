/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

// 由于 Map 的 Zero value 是 nil，因此 nil Map 没有 Key，且不能添加 Key。
var m map[string]Vertex

func main() {
	// 验证 Map 的 Zero value
	fmt.Println(m)
	if m == nil {
		fmt.Println("nil Map")
	}

	// make function 返回给定 Type 的 Map，且已初始化可供使用。
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68433, -74.39967,
	}
	fmt.Println(m["Bell Labs"])
}
