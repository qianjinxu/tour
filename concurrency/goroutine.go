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

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(1000 * time.Millisecond) // 暂停当前 goroutine 至少持续时间 1000ms。如果持续时间是零或负数，则其立即返回。
		fmt.Println(time.Now())
		fmt.Printf("%v\n", s)
	}
}

func main() {
	go say("Qianjin Xu") // 启动新 goroutine 以运行 sqy function
	say("Hello")
}
