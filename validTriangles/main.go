package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var t, a, b, c int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	r := strings.NewReader(scanner.Text())
	fmt.Fscanf(r, "%d", &t)

	for i := 0; i < t; i++ {

		scanner.Scan()
		r = strings.NewReader(scanner.Text())
		fmt.Fscanf(r, "%d %d %d", &a, &b, &c)
		if (a + b + c) == 180 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}
