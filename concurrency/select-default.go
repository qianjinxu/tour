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
	tick := time.Tick(2 * time.Second)
	boom := time.After(5 * time.Second)
	for {
		select {
		// 该 case 需要 2s 准备好
		case <-tick:
			fmt.Println("tick.")
		// 该 case 需要 5s 准备好
		case <-boom:
			fmt.Println("boom!")
			return
		// 默认 case（如果其它 case 没有准备好，则运行默认 case）
		default:
			fmt.Println(time.Now())
			time.Sleep(1 * time.Second)
		}
	}
}
