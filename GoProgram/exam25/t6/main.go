package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		n     int
		count int
		win   int
	)
	fmt.Scan(&n)
	Apower := make([]int, n)
	Bpower := make([]int, n)
	for i := range Apower {
		fmt.Scan(&Apower[i])
	}
	for i := range Bpower {
		fmt.Scan(&Bpower[i])
	}
	sort.Slice(Apower, func(i, j int) bool {
		return Apower[i] < Apower[j]
	})
	sort.Slice(Bpower, func(i, j int) bool {
		return Bpower[i] < Bpower[j]
	})
	// fmt.Println(Apower)
	// fmt.Println(Bpower)
	for i := 0; i < n; i++ {
		if Apower[i] > Bpower[0] {
			count = i
			break
		}
	}
	for j := count; j < n; j++ {
		if Apower[j] > Bpower[j-count] {
			win += 1
		} else {
			break
		}
	}
	win = win*200 - (n-win)*200
	fmt.Println(win)
}
