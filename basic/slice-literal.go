/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	q := []int{
		2,
		14,
		49,
		53,
		98,
	}
	fmt.Println(q)

	// Slice literal
	r := []bool{
		true,
		false,
		true,
		false,
		true,
	}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{14, false},
		{49, true},
		{53, false},
		{98, true},
	}
	fmt.Println(s)
}
