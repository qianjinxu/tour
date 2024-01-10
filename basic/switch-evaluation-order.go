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
	// fmt.Println(time.Now().Weekday())
	fmt.Println("When's Saturday?")

	today := time.Now().Weekday()
	// switch statement 从上往下评估每个 case，如果 case 评估成功，则停止评估。
	switch time.Saturday {
	case today + 0:
		fmt.Println("Today.")
	case today + 1:
		fmt.Println("Tomorrow.")
	case today + 2:
		fmt.Println("In two days.")
	default:
		fmt.Println("Too far away.")
	}
}
