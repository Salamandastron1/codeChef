package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numLines := 0
	for _, v := range scanner.Bytes() {
		numLines = (numLines * 10) + int(rune(v)-'0')
	}
	fmt.Println("num lines", numLines)
	for i := 0; i < numLines; i++ {
		var count int
		var n int
		scanner.Scan()

		for _, v := range scanner.Bytes() {
			n = (n * 10) + int(rune(v)-'0')
		}

		for n > 0 {
			r := n % 10
			n = n / 10
			if r == 4 {
				count++
			}
		}
		fmt.Println(count)
	}
}
