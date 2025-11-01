package main

import (
	"fmt"
	"strings"
)

func main() {
	var n1 int
	fmt.Scan(&n1)
	answer := make([]string, n1)
	for j := 0; j < n1; j++ {
		var n int
		fmt.Scan(&n)

		nums := make([]int, n)
		for i := range nums {
			fmt.Scan(&nums[i])
		}
		var win bool
		if n%2 != 0 {
			win = true
		}
		if win {
			answer[j] = "syj\n"
		} else {
			answer[j] = "cc\n"
		}
	}
	Answer := strings.Join(answer, "")
	fmt.Println(Answer)
}

// 获胜条件：
// 1，交换后形成0101循环，长度为偶
// 2，取走最后一个数，长度为0
// 若起始长度为奇数，先手无论哪个回合都必定能取数，syj胜
// 若起始长度为偶数，则先手取完数后剩余数量必定为奇数，cc胜
