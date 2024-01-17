/*
 * GitHub: github.com/qianjinxu
 * Email: xuqianjinchn@gmail.com
 * Bio: https://jin.bio
 */

// 任务说明：https://go.dev/tour/methods/18
package main

import "fmt"

type IPAddr [4]byte

// 为 IPAddr type 声明一个 String() method 以实现 Stringer interface
func (i IPAddr) String() string {
	return fmt.Sprintf("%d.%d.%d.%d", i[0], i[1], i[2], i[3]) // 将 IP 地址打印为点分十进制格式
}

func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}
