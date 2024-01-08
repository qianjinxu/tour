/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import (
	"fmt"
	"time"
)

func main() {
	// fmt.Println(time.Now().Hour())

	t := time.Now()
	// 不带条件的 switch statement 与 switch true 相同，这种结构是编写较长 if-then-else 链的简洁方法。
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good eventing.")
	}
}
