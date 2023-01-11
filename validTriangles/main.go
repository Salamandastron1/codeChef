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
	numLines, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

	for i := 0; i < numLines; i++ {
		var sumAngles int

		scanner.Scan()
		for _, v := range strings.Split(scanner.Text(), " ") {
			num, _ := strconv.Atoi(v)
			sumAngles += num
		}
		if sumAngles == 180 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
