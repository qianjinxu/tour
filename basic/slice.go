/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	// Array 大小是固定的
	primes := [6]int{
		4,
		12,
		7,
		22,
		31,
		19,
	}
	fmt.Println(primes)

	var (
		// Slice 大小是动态的
		// 将 s variable 声明为一个包含 int type 的 Slice，该 Slice element 通过指定 2 个 Indices (lowBound, highBound) 形成的，并使用 : 分隔。这意味着将选择一个 half-open range，其中，不包括第 1 个 Element，但包括最后 1 个 Element。
		s []int = primes[1:4]
	)
	fmt.Println(s)
}
