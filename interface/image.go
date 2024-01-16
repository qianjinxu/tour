/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import (
	"fmt"
	"image"
)

// type Image interface {
// 	// color.Color type 和 color.Model type 也是 Interface，但将通过使用预定义的实现 color.RGBA 和 color.RGBAModel 来忽略它。
// 	ColorModel() color.Model
// 	At(x, y int) color.Color
// 	// Bounds() method 的 Rectangle 返回值实际是一个 image.Rectangle type，因为声明位于 image package 内部。
// 	Bounds() Rectangle
// }

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
