package main

import (
	"fmt"
)

func main() {
	var (
		x    int
		y    int
		m    int
		n    int
		L    int
		time int
	)
	fmt.Scan(&x, &y, &m, &n, &L)
	if n == m && x != y {
		fmt.Println("Impossible")
	} else {
		for x != y {
			x += m
			y += n
			if x >= 5 {
				x = x - 5
			}
			if y >= 5 {
				y = y - 5
			}
			time += 1
		}
		fmt.Println(time)
	}
}
