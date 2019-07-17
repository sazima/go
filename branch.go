package main

import (
	"fmt"
	"io/ioutil"
)

func grade(score int) string {
	g := ""
	switch {
	case score < 60:
		g = "f"
	case score < 70:
		g = "e"
	case score < 80:
		g = "d"
	case score < 90:
		g = "c"
	default:
		panic(fmt.Sprintf("wrong score %d", score))
	}

	return g
}

func main() {
	const filename = "abc.txt"
	if contents, err := ioutil.ReadFile(filename); err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%s \n ", contents)
	}
	s := grade(11)
	fmt.Println(s)
	switch s {
	case "f":
		fmt.Printf("terrible")
	case "a":
		fmt.Print("great")
	default:
		fmt.Printf("I don't konw")

	}

}
