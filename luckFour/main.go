package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var t int
	fmt.Scanln(&t)
	for t > 0 {
		var n int
		fmt.Scanln(&n)
		fmt.Println(strings.Count(strconv.Itoa(n), "4"))
		t--
	}
}
