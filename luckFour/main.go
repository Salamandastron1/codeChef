package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numLines, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	for i := 0; i < numLines; i++ {
		scanner.Scan()

		var count int
		for _, v := range scanner.Text() {
			if v == '4' {
				count++
			}
		}
		fmt.Println(count)
	}
}
