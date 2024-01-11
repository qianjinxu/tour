/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	// 将 a variable 声明为一个包含 2 个 string type 的 Array，该 Array 包含 2 个 Element。
	var a [2]string
	a[0] = "Hello,"
	a[1] = "Qianjin Xu!"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	// 由于 Array length 是 [2]string type 的一部分，因此无法调整 Array 大小。
	primes := [6]int{
		2,
		3,
		5,
		7,
		11,
		13,
	}
	fmt.Println(primes)
}
