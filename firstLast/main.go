package main

import "fmt"

func main() {
	var t, n int
	_, err := fmt.Scan(&t)
	if err != nil {
		panic(err)
	}
	fmt.Println(t)
	for i := 0; i < t; i++ {
		_, err = fmt.Scan(&n)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%d\n", sum(n))
	}
}

func sum(n int) int {
	var first int
	last := n % 10
	for n != 0 {
		first = n % 10
		n = n / 10
	}
	return first + last
}
