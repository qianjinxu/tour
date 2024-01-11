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

// Map literal 与 Struct literal 类似，但 Key 是必需的。
var m = map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
	"Tesla": Vertex{},
}

func main() {
	fmt.Println(m)
}
