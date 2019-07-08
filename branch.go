package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	const filename = "abc.txt"
	contents, error := ioutil.ReadFile(filename)
	if error != nil {
        fmt.Println(error)
	}else{
		fmt.Printf("%s \n", contents)
	}

}
