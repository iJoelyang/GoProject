package main

import (
	"fmt"
	"strings"
)

func lengthOfNonRepeatingSubStr(s string) int {
	lastOccurred := make(map[rune]int) // empty map
	start := 0
	maxLength := 0

	for i, ch := range []rune(s) {
		lastI, ok := lastOccurred[ch]
		if ok && lastI >= start {
			start = lastI + 1
		}
		if i-start+1 > maxLength {
			maxLength = i - start + 1
		}
		lastOccurred[ch] = i
	}
	return maxLength
}

func main() {
	fmt.Println(lengthOfNonRepeatingSubStr("abcda"))
	fmt.Println(lengthOfNonRepeatingSubStr("a"))
	fmt.Println(lengthOfNonRepeatingSubStr("pwwkew"))
	// 使用byte，不能正确处理中文，需要转成rune
	fmt.Println(lengthOfNonRepeatingSubStr("大家好"))
	fmt.Println(lengthOfNonRepeatingSubStr("三二四三二一"))

	// 其他字符串操作
	lists := strings.Fields("aaaa  ss d")
	fmt.Println("lists: ", lists)
}
