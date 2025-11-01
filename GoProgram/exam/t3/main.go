package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

// Word 结构体用于存储每个单词的信息
type Word struct {
	word        string // 原始单词
	tailChar    byte   // 尾部字符
	repeatCount int    // 尾部字符重复次数
	length      int    // 单词长度
	dictionary  string // 字典序（与原始单词相同）
}

// 实现 sort.Interface 接口以便对单词进行排序
type ByCriteria []Word

func (a ByCriteria) Len() int      { return len(a) }
func (a ByCriteria) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByCriteria) Less(i, j int) bool {
	// 根据题目要求的规则进行比较
	if a[i].repeatCount != a[j].repeatCount {
		return a[i].repeatCount > a[j].repeatCount // 尾部字符重复次数越大越好
	}
	if a[i].tailChar != a[j].tailChar {
		return a[i].tailChar < a[j].tailChar // 字典序越小越好
	}
	if a[i].length != a[j].length {
		return a[i].length < a[j].length // 长度越小越好
	}
	return a[i].dictionary < a[j].dictionary // 字典序越小越好
}

// findBestWord 函数接收句子并返回最好的“好词”
func findBestWord(sentence string) string {
	// 将句子转换为小写并分割成单词
	words := strings.Fields(strings.ToLower(sentence))
	// 展开缩写
	words = expandAbbreviations(words)

	var wordList []Word

	// 遍历每个单词，计算其尾部信息并存储在 wordList 中
	for _, word := range words {
		tailChar, repeatCount := findTailInfo(word)
		currentWord := Word{
			word:        word,
			tailChar:    tailChar,
			repeatCount: repeatCount,
			length:      len(word),
			dictionary:  word,
		}
		wordList = append(wordList, currentWord)
	}

	// 对单词列表进行排序
	sort.Sort(ByCriteria(wordList))

	// 返回排序后第一个单词作为最好的“好词”
	if len(wordList) > 0 {
		return wordList[0].word
	}
	return ""
}

// expandAbbreviations 函数将缩写还原为完整单词
func expandAbbreviations(words []string) []string {
	var expandedWords []string
	for _, word := range words {
		switch word {
		case "i'll":
			expandedWords = append(expandedWords, "i", "will")
		case "you'll":
			expandedWords = append(expandedWords, "you", "will")
		case "I'm":
			expandedWords = append(expandedWords, "i", "am")
		default:
			expandedWords = append(expandedWords, word)
		}
	}
	return expandedWords
}

// findTailInfo 函数计算单词的尾部字符及其重复次数
func findTailInfo(word string) (byte, int) {
	n := len(word)
	if n == 0 {
		return 0, 0
	}

	tailChar := word[n-1]
	repeatCount := 1

	// 计算尾部字符的重复次数
	for i := n - 2; i >= 0 && word[i] == tailChar; i-- {
		repeatCount++
	}

	return tailChar, repeatCount
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("请输入一个英语长句:")
	sentence, _ := reader.ReadString('\n') // 读取整行输入
	sentence = strings.TrimSpace(sentence) // 去除前后空格

	// 找到最好的“好词”
	bestWord := findBestWord(sentence)
	fmt.Println("最好的好词是:", bestWord)
}
