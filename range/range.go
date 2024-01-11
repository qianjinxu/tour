/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

var car = []int{1, 2, 3}

func main() {
	// for loop 的 Range 在 Slice 或 Map 上进行迭代
	// 如果 Range 覆盖 Slice，则每次迭代会返回 2 个值：第 1 个是 index，第 2 个是该 index 处 Slice element 的副本。
	for i, v := range car {
		fmt.Printf("index=%d, sliceElement=[%d]\n", i, v)
	}
}
