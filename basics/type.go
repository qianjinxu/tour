/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import (
	"fmt"
	"math/cmplx"
)

var (
	ToBe bool = false
	// int, uint, uintptr type 在 32-bit 操作系统上为 32 bit 宽度，在 64-bit 操作系统上为 64 bit 宽度。
	// 如果需要整数值，则应使用 int type。
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	fmt.Printf("Type: %T, Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T, Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T, Value: %v\n", z, z)
}
