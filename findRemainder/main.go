package main

import "fmt"

func main() {
	var t, d, q int
	fmt.Scanf("%d", &t)
	for ; t > 0; t-- {
		fmt.Scanf("%d %d", &d, &q)
		fmt.Println(d % q)
	}
}
