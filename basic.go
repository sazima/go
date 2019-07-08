package main

import (
	"fmt"
	"math"
)

var aa = true

var (
	bb = false
	ss = "string "
)

func variableZeroValue() {
	var a int
	var s string
	fmt.Printf("%d, %q\n", a, s)
}

func variableInitialValue() {
	var a, c int = 3, 4
	var b string = "abc"
	fmt.Println(a, b, c)
}

func varialbeTypeDeuction() {
	a, b, s := 3, 4, "helloworld"
	fmt.Println(a, b, s)
}

func consts() {
	const filename = "abc.txt"
	const a, b = 3.2, 4
	ccc := math.Sqrt(a*a + b*b)
	fmt.Println(a, b, ccc)
}

func emums() {
	const (
		cpp = iota
		java
		python
		golang
	)
	fmt.Println(cpp, java, python, golang)
	const (
		b = 1 << (10 * iota)
		kb
		mb
		tb
		pb
	)
	fmt.Println(b, kb, mb, tb, pb)

}
func main() {
	fmt.Println("hello world")
	variableZeroValue()
	variableInitialValue()
	varialbeTypeDeuction()
	fmt.Println(aa)
	fmt.Println(bb, ss)
	consts()
	emums()
}
