package main

import (
	"fmt"
	"strings"
)

const QUEUE_SIZE = 8

type ParseResult struct {
	CurrentQueue string
	Output       []string
}

var skillPatterns1 = map[string]string{
	"SDJKL": "终极波动拳",
}
var skillPatterns2 = map[string]string{
	"SDJK":  "波动连击",
	"DSDJJ": "升龙连破",
	"ASDLK": "反击爆破",
	"SDJKL": "终极波动拳",
}

var skillPatterns3 = map[string]string{
	"SDJ":  "波动拳",
	"DSDJ": "升龙拳",
	"ASDL": "防御反击",
	"WASJ": "旋风拳",
	"DJJJ": "三连击",
}

func parseCommand1(input string) ParseResult {
	result := ParseResult{
		CurrentQueue: "",
		Output:       []string{},
	}
	queue := []rune{}
	for _, char := range input {
		if char == 'T' {
			queue = []rune{}
			continue
		}
		if len(queue) > 0 && (char == 'W' || char == 'A' || char == 'S' || char == 'D') && queue[len(queue)-1] == char {
			continue
		}
		queue = append(queue, char)
		if len(queue) > QUEUE_SIZE {
			queue = queue[1:]
		}
		for pattern, skill := range skillPatterns1 {
			if strings.HasPrefix(string(queue), pattern) {
				result.Output = append(result.Output, skill)
				queue = queue[len(pattern):]
				break
			}
		}
	}
	result.CurrentQueue = string(queue)
	return result
}
func parseCommand2(input string) ParseResult {
	result := ParseResult{
		CurrentQueue: "",
		Output:       []string{},
	}
	queue := []rune{}
	for _, char := range input {
		if char == 'T' {
			queue = []rune{}
			continue
		}
		if len(queue) > 0 && (char == 'W' || char == 'A' || char == 'S' || char == 'D') && queue[len(queue)-1] == char {
			continue
		}
		queue = append(queue, char)
		if len(queue) > QUEUE_SIZE {
			queue = queue[1:]
		}
		for pattern, skill := range skillPatterns2 {
			if strings.HasPrefix(string(queue), pattern) {
				result.Output = append(result.Output, skill)
				queue = queue[len(pattern):]
				break
			}
		}
	}
	result.CurrentQueue = string(queue)
	return result
}
func parseCommand3(input string) ParseResult {
	result := ParseResult{
		CurrentQueue: "",
		Output:       []string{},
	}
	queue := []rune{}
	for _, char := range input {
		if char == 'T' {
			queue = []rune{}
			continue
		}
		if len(queue) > 0 && (char == 'W' || char == 'A' || char == 'S' || char == 'D') && queue[len(queue)-1] == char {
			continue
		}
		queue = append(queue, char)
		if len(queue) > QUEUE_SIZE {
			queue = queue[1:]
		}
		for pattern, skill := range skillPatterns3 {
			if strings.HasPrefix(string(queue), pattern) {
				result.Output = append(result.Output, skill)
				queue = queue[len(pattern):]
				break
			}
		}
	}
	result.CurrentQueue = string(queue)
	return result
}
func main() {
	var input string
	fmt.Scanln(&input)

	result := parseCommand1(input)
	if len(result.Output) == 0 {
		result = parseCommand2(input)
	}
	if len(result.Output) == 0 {
		result = parseCommand3(input)
	}
	fmt.Printf("ParseResult{\n")
	fmt.Printf("    CurrentQueue: \"%s\",\n", result.CurrentQueue)
	fmt.Printf("    Output:       %v,\n", result.Output)
	fmt.Printf("}\n")
}

//go run main.go
