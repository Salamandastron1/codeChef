package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var x, y int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numLines, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	if err != nil {
		panic(err)
	}

	for i := 0; i < numLines; i++ {
		scanner.Scan()
		input := strings.Split(scanner.Text(), " ")
		x, _ = strconv.Atoi(input[1])
		y, _ = strconv.Atoi(input[2])
		if x > y {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
