package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	s := "User input: \n" + strings.Join(os.Args[1:], "\n")
	fmt.Println(s)
}
