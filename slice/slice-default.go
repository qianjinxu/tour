/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	b := []int{1, 2, 3, 4, 5}
	c := []int{1, 2, 3, 4, 5}
	d := []int{1, 2, 3, 4, 5}

	// 如果未省略 lowBound 和 upBound，则使用它们的设定值。
	a = a[1:5]
	fmt.Println(a)

	// 如果省略 lowBound，则使用其默认值（0）
	b = b[:5]
	fmt.Println(b)

	// 如果省略 upBound，则使用其默认值（Slice length）
	c = c[1:]
	fmt.Println(c)

	// 如果省略 lowBound 和 upBound，则使用它们的默认值。
	d = d[:]
	fmt.Println(d)
}
