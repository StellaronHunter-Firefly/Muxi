package main

import (
	"fmt" // 导入 fmt 包，用于格式化输入输出
)

func jump(nums []int) bool {
	var ok bool // 定义布尔变量 ok，用于存储最终结果，默认为 false

	n := len(nums) // 获取数组 nums 的长度，即格子的数量
	if n == 0 {
		return ok // 如果格子数量为 0，直接返回 false
	}

	maxReach := 0 // 初始化 maxReach 变量，表示当前能到达的最远位置

	for i := 0; i < n; i++ { // 遍历每个格子
		if i > maxReach { // 如果当前位置超过了最大可达范围，则无法继续前进
			return false
		}
		maxReach = max(maxReach, i+nums[i]) // 更新最大可达范围
		if maxReach >= n-1 {                // 如果最大可达范围已经大于等于最后一个格子的位置，则可以到达终点
			return true
		}
	}

	return ok // 返回最终结果，如果没有达到终点则返回 false
}

func max(a, b int) int { // 辅助函数，用于比较两个整数并返回较大的那个
	if a > b {
		return a
	}
	return b
}

func main() {
	var n int
	fmt.Scan(&n) // 读取第一行输入，即格子的数量，并存储在变量 n 中

	nums := make([]int, n) // 创建一个长度为 n 的整数切片 nums
	for i := range nums {  // 遍历 nums 切片
		fmt.Scan(&nums[i]) // 读取第二行输入中的每个整数，并存储在 nums 切片中
	}

	if jump(nums) { // 调用 jump 函数，传入 nums 切片
		fmt.Println("Yes") // 如果 jump 函数返回 true，则输出 "Yes"
	} else {
		fmt.Println("No") // 否则输出 "No"
	}
}
