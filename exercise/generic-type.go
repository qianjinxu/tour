/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

// 任务说明：https://go.dev/tour/generics/2
package main

type Mygeneric[T any] struct {
	next *Mygeneric[T]
	val  T
}

func main() {

}
