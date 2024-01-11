/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	car := make([]int, 10)
	fmt.Println(car)

	for i := range car {
		car[i] = 1 << uint(i) // 等于 2**i
		fmt.Println(car)
	}

	for index, sliceElement := range car {
		fmt.Printf("index=%d, sliceElement=[%d]\n", index, sliceElement)
	}

	// 使用 _ 忽略 index
	// for _, sliceElement := range car {
	// 	fmt.Printf("sliceElement=[%d]\n", sliceElement)
	// }

	// 使用 _ 忽略 sliceElement
	// for index, _ := range car {
	// 	fmt.Printf("index=%d\n", index)
	// }
}
