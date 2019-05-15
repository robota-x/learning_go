package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	var s, sep string

	sep = " "
	s = "User input:"

	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
	}

	fmt.Println(s)

	sep = "\n"
	s = "User input: \n"

	for i, arg := range os.Args[1:] {
		s += strconv.Itoa(i) + ")" + arg + sep
	}

	fmt.Println(s)
}
