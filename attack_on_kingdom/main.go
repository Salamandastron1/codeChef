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
		// don't actually need number of elements
		// so scan past
		s.Scan()
		temps := strings.Split(s.Text(), " ")
		sort.Slice(temps, func(i, j int) bool {
			return parseInt(temps[i]) < parseInt(temps[j])
		})
		fmt.Println(parseInt(temps[1]))
	}
}

func parseInt(s string) int {
	num, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return num
}
