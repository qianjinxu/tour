/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import "fmt"

func main() {
	// for statement 由 init statement, condition expression, post statement 等 3 部分组成，各部分之间使用分号分割。
	// init statement 在第一次迭代之前执行，它通常是一个 short variable 声明，且该 short variable 仅在 for statement 范围内可见。
	// condition expression 即条件表达式，在每次迭代之前评估，如果 boolean 计算结果为 false，则终止迭代。
	// post statement 在每次迭代结束时执行
	sum := 0
	for i := 0; i < 9; i++ {
		sum += i
	}

	fmt.Println(sum)
}
