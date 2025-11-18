package main

import (
	"fmt"
)

func main() {
	s := make([]byte, 5)
	fmt.Printf("len:%v, cap:%v\n", len(s), cap(s))
	slices1 := "hello，世界"
	var i int
	for k := range slices1 {
		i = k
	}

	fmt.Printf("长度：%v, 循环次数：%v\n", len(slices1), i)
	if len(slices1) > i {
		fmt.Println("长度大于循环次数")
	} else if len(slices1) < i {
		fmt.Println("长度大于循环次数")
	} else {
		fmt.Println("长度等于循环次数")
	}
}
