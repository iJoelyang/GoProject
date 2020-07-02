package main

import (
	"fmt"
	"math"
	"math/cmplx"
)

var (
	aa = 3
	bb = true
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d %q\n", a, s)
}

func variableInintialValue() {
	var a, b = 3, 4
	var s = "abc"
	fmt.Println(a, b, s)
}

func variableTypeDeduction() {
	var a, b, c, s, bl = 1, 2, 3, "abc", true
	//a, b, c, s, bl := 1, 2, 3, "abc", true
	b = 5
	fmt.Println(a, b, c, s, bl)
}

func euler() {
	c := 3 + 4i
	fmt.Printf("\n%v\n", cmplx.Abs(c))
	fmt.Printf("%.3f\n", cmplx.Pow(math.E, 1i*math.Pi)+1)
}

func triangle() {
	var a, b = 3, 4
	var c int
	c = int(math.Sqrt(float64(a*a + b*b)))
	fmt.Println(c)
}

func consts() {
	const a, b = 3, 4
	const filename = "abc.txt"
	var c int
	c = int(math.Sqrt(a*a + b*b))
	fmt.Println(c)
}

func enums() {
	const (
		cpp = iota
		_
		python
		java
	)

	const (
		b = 1 << (10 + iota)
		//b = 1 << iota
		kb
		mb
		gb
		tb
		pb
	)
	fmt.Println(cpp, python, java)
	fmt.Println(b, kb, mb, gb, tb, pb)
}

func main() {
	fmt.Println("Hello World!")
	variableZeroValue()
	variableInintialValue()
	variableTypeDeduction()
	fmt.Printf("aa = %d, bb = %T", aa, bb)
	euler()
	triangle()
	consts()
	enums()
}
