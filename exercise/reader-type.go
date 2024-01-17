/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

// 任务说明：https://go.dev/tour/methods/22
package main

import (
	"golang.org/x/tour/reader"
)

type Myreader struct {
}

// 实现一个发出无限 ASCII 字符数据流的 Reader type
func (m Myreader) Read(bytes []byte) (int, error) {
	for i := range bytes {
		bytes[i] = 65
	}
	return len(bytes), nil
}

func main() {
	reader.Validate(Myreader{})
}
