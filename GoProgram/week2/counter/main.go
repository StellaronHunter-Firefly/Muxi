package main

import (
	"fmt"
	"time"
)

func Counter() func() int {
	cot := 0
	return func() int {
		cot++
		return cot
	}
}

func main() {
	counter := Counter()
	var a int
	fmt.Println("请输入需要计数的次数：")
	fmt.Scanf("%d", &a)
	for i := 0; i < a; i++ {
		fmt.Println(counter())
		time.Sleep(500 * time.Millisecond)
	}
}
