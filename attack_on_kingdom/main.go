package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// an iterator counter for later
	s := bufio.NewScanner(os.Stdin)
	// don't need the num lines
	// will rely on scanner tokens
	s.Scan()

	for s.Scan() {
		n := parseInt(s.Text())
		arr := make([]int, n, n)
		s.Scan()
		temps := strings.Split(s.Text(), " ")
		for i := 0; i < n; i++ {
			arr[i] = parseInt(temps[i])
		}

		sort.Ints(arr)
		fmt.Println(arr[1])
	}
}

func parseInt(s string) int {
	num, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return num
}
