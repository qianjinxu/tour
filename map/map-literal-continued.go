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

// 如果 top-level Type 只是 Type 名称，则可以从 Map literal element 中省略它。
var m = map[string]Vertex{
	"Bell Labs": {40.68433, -74.39967},
	"Google":    {37.42202, -122.08408},
	"Tesla":     {},
}

func main() {
	fmt.Println(m)
}
