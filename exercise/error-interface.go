/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

// 任务说明：https://go.dev/tour/methods/20
package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("Cannot Sqrt negative number: %v", float64(e)) // 将 e 转换为 float64(e) 以避免使 Go 程序陷入死循环
}

// 该 Function 被设计为只对正数求平方根。如果给定负数，则返回错误。
func Sqrt(x float64) (float64, error) {
	// 如果给定负数，则返回 ErrNegativeSqrt value。因为该 Function 不打算支持负数。
	if x < 0 {
		return 0, ErrNegativeSqrt(x)
	}
	// 如果给定正数，则返回平方根。
	z := 1.0
	for i := 1; i <= 100; i++ {
		z -= (z*z - x) / (2 * z)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(25))
	fmt.Println(Sqrt(9))
	fmt.Println(Sqrt(-2))
}
