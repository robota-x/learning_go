package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	countMap := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)

	for input.Scan() {
		countMap[input.Text()]++
	}

	fmt.Printf("Word\t\tCount\n")
	for line, n := range countMap {
		if n > 0 {
			fmt.Printf("%d\t\t%s\n", n, line)
		}
	}
}
