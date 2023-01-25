package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for scanner.Scan() {
		num, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
		count := 0
		for num%2 == 0 {
			num /= 2
			count++
		}
		if count%2 == 0 {
			fmt.Println(1)
		} else {
			fmt.Println(0)
		}
	}
}
