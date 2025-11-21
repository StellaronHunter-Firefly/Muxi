package main

import (
	"fmt"
	"sync"
)

func main() {
	// 定义两个channel，一个打印数字，一个打印字母
	numCh := make(chan struct{})
	strCh := make(chan struct{})
	var wg sync.WaitGroup
	wg.Add(2)
	// 打印字母
	go func() {
		defer wg.Done()
		for i := 'A'; i <= 'Z'; i += 2 {
			fmt.Printf(string(i))
			fmt.Printf(string(i + 1))
			// 通知打印数字
			numCh <- struct{}{}
			// 阻塞等待打印字母
			<-strCh
		}
	}()
	// 打印数字
	go func() {
		defer wg.Done()
		for i := 0; i < 26; i += 2 {
			<-numCh
			fmt.Printf("%v", i)
			fmt.Printf("%v", i+1)
			// 通知打印字母
			strCh <- struct{}{}
		}
	}()
	wg.Wait()
	fmt.Println("\nfinished")
}
