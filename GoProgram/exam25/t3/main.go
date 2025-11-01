package main

import (
	"fmt"
	"strings"
)

func main() {
	var (
		n        int
		E        int
		r        int
		location int = 1
		s        string
		answer   int = 1
		count1   int
		count2   int
	)

	fmt.Scan(&n)
	fmt.Scan(&E)
	fmt.Scan(&r)
	fmt.Scan(&s)
	nums := strings.Split(s, "")
	for i := range nums {
		if nums[i] == "+" {
			count1 += 1
		}

	}
	for g := n - 2; g > 0; g-- {
		if nums[g] == "-" {
			count2 += 1
		} else {
			break
		}
	}
	if nums[n-1] == "-" || E+r*count1 < count2 {
		answer = -1
	} else {
		for location+E < n {
			for g := location - 1; g < n; g++ {
				count2 = 0
				if nums[g] == "-" {
					count2 += 1
				} else {
					break
				}
			}
			if count2 >= E {
				answer = -1
				break
			}
			if nums[location+E-1] != "-" {
				if n-location <= 2*E {
					location += E
					answer += 1
				} else {
					for j := location + E; j > location; j-- {
						if nums[j-1] == "+" {
							location += j - location
							answer += 1
							break
						}
					}
				}
			} else {
				for j := location + E; j > location; j-- {
					if nums[j-1] != "-" {
						location += j - location
						answer += 1
						break
					}
				}
			}
			fmt.Println(location, E, r)
		}

	}

	fmt.Println(answer)
}

//cd D:\GoProgram\exam25\t3
