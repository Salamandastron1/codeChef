package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var m any
	var x int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for scanner.Scan() {
		x = toInt(scanner.Bytes())
		num := x + 1
		for !unique(num) {
			num++
		}
		fmt.Println(num)
	}
}
func unique(i int) bool {
	memo := map[int]bool{}

	for i > 0 {
		digit := i % 10
		if ok := memo[digit]; ok {
			return false
		}

		memo[digit] = true
		i /= 10
	}
	return true
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}
