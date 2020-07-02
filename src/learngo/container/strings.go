package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	s := "我爱北京天安门"
	fmt.Println(s)
	bytes := []byte(s)
	fmt.Printf("%X\n", )

	for i, b := range bytes {
		fmt.Printf("(%d %X) ", i, b)
	}

	fmt.Println()
	fmt.Println("Rune count: ", utf8.RuneCountInString(s))

	for len(bytes) > 0 {
		ch, size := utf8.DecodeRune(bytes)
		bytes = bytes[size:]
		fmt.Printf("%c ", ch)
	}

	fmt.Println()
	for i, ch := range []rune(s) {
		fmt.Printf("(%d, %c) ", i, ch)
	}
}
