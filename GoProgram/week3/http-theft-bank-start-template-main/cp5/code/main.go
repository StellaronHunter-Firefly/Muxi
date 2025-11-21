/*
>示例如下：
============================================
输入：
3
1 2 3
输出：
[[1 2 3][1 3 2][2 1 3][2 3 1][3 1 2][3 2 1]]
============================================

>代码模板:
*/
package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var res [][]int
	var backtrack func(path []int, used []bool)
	backtrack = func(path []int, used []bool) {
		if len(path) == len(nums) {
			tmp := make([]int, len(path))
			copy(tmp, path)
			res = append(res, tmp)
			return
		}
		for i := 0; i < len(nums); i++ {
			if used[i] {
				continue
			}
			used[i] = true
			path = append(path, nums[i])
			backtrack(path, used)
			used[i] = false
			path = path[:len(path)-1]
		}
	}
	used := make([]bool, len(nums))
	backtrack([]int{}, used)
	return res
}

func main() {
	var n, a int
	fmt.Scanf("%d", &n)

	testSlice := make([]int, n)
	// 标准输入n个不重复的数字
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &a)
		testSlice[i] = a + 1
	}
	res := permute(testSlice)
	fmt.Println(res)
}
