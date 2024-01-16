/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import (
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("Hello, Qianjin Xu!")

	b := make([]byte, 8)

	for {
		// Read method 使用给定的数据填充 byte slice
		n, err := r.Read(b)

		// Read method 返回填充的 byte 数和  error value
		fmt.Printf("n = %v err = %v b= %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		// 当数据流结束时，则返回 io.EOF 错误。
		if err == io.EOF {
			break
		}
	}
}
