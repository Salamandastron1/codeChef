package main

import "fmt"

func main() {

	var t int
	fmt.Scan(&t)

	//println(t)

	for t > 0 {
		var num int
		fmt.Scan(&num)

		var sum int = 0

		for num > 0 {
			fmt.Println(num)
			sum = sum + num%10
			num = num / 10
		}

		fmt.Println(sum)

		t--
	}

}
