package main

import (
	"fmt"
)

func main() {
	var k int
	fmt.Scan(&k)
	hot := make([]int, k)
	for i := range hot {
		fmt.Scan(&hot[i])
	}
	var answer int = -1
	for j := 0; j < k-2; j++ {
		if hot[j]%2 == 0 && hot[j+1]%2 != 0 && hot[j+2]%2 == 0 && answer < hot[j]+hot[j+1]+hot[j+2] {
			answer = hot[j] + hot[j+1] + hot[j+2]
		}
	}
	fmt.Println(answer)
}
