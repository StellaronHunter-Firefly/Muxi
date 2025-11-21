package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"time"
)

var wg sync.WaitGroup

type GoroutineResult struct {
	ID        int
	SleepTime int
}

func fn(i int, ch1 chan GoroutineResult) {
	defer wg.Done()
	randomNumber := rand.Intn(1000) + 1
	time.Sleep(time.Duration(randomNumber) * time.Millisecond)
	ch1 <- GoroutineResult{ID: i, SleepTime: randomNumber}
	fmt.Printf("Goroutine %d 休眠了 %d 毫秒\n", i, randomNumber)
}

func main() {
	ch1 := make(chan GoroutineResult, 20)
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go fn(i, ch1)
	}
	wg.Wait()
	close(ch1)
	var results []GoroutineResult
	for result := range ch1 {
		results = append(results, result)
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].ID < results[j].ID
	})
	fmt.Println("所有 Goroutine 按编号排序结果:")
	for _, result := range results {
		fmt.Printf("Goroutine %d 休眠了 %d 毫秒\n", result.ID, result.SleepTime)
	}
}
