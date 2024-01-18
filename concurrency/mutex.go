/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

package main

import (
	"fmt"
	"sync"
	"time"
)

// SafeCounter mutex 可以安全使用 Concurrency
type SafeCounter struct {
	mu sync.Mutex // sync.Mutex 及其 Lock(), Unlock() method 提供了 Mutex
	v  map[string]int
}

// Inc method 用于增量给定 key 的计数器
func (c *SafeCounter) Inc(key string) {
	// 调用 Lock() method 以实现每次只有一个 goroutine 可以访问 c.v map
	c.mu.Lock()
	c.v[key]++
	c.mu.Unlock()
}

// Value method 用于返回给定 key 的计数器的当前 Value
func (c *SafeCounter) Value(key string) int {
	// 调用 Lock() method 以实现每次只有一个 goroutine 可以访问 c.v map
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("hello")
	}
	time.Sleep(2 * time.Second)
	fmt.Println(c.Value("hello"))
}
