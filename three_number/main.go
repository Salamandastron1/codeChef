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
		xi := make([]int, 3, 3)
		temps := strings.Split(scanner.Text(), " ")
		for i, v := range temps {
			num, _ := strconv.Atoi(v)
			xi[i] = num
		}
		if xi[0]+xi[1] == xi[2] || xi[1]+xi[2] == xi[0] || xi[0]+xi[2] == xi[1] {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func parseInt(s string) int {
	num, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return num
}
