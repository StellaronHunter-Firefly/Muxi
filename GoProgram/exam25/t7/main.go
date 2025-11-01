package main

import (
	"bufio"
	"fmt"
	"net/http"
	//	"net/url"
	"os"
)

func main() {
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("ListenAndServe: ", err)
	}
	file, err := os.Open("text.txt")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
}
