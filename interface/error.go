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

type Errorlist struct {
	When time.Time
	What string
}

// 为 *Errorlist pointer 声明一个 Error() method 以实现 error interface
func (e *Errorlist) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &Errorlist{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	// 通过 Function 调用时返回 Error value 判断是否成功
	err := run()
	if err != nil {
		fmt.Println(err)
	}
}
