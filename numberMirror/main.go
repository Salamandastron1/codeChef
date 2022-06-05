// We have populated the solutions for the 10 easiest problems for your support.
// Click on the SUBMIT button to make a submission to this problem.

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, k int
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Scanln(&n, &k)
	counter := 0
	for scanner.Scan() {
		v, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
		}
		if int(v)%k == 0 {
			counter += 1
		}
	}
	fmt.Println(counter)
}
