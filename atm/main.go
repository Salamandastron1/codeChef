// We have populated the solutions for the 10 easiest problems for your support.
// Click on the SUBMIT button to make a submission to this problem.

package main

import "fmt"

func main() {

	var withdrawlAmt int
	fmt.Scan(&withdrawlAmt)

	var currBalance float64
	fmt.Scan(&currBalance)

	if withdrawlAmt%5 == 0 && (currBalance-0.5-float64(withdrawlAmt) >= 0) {
		fmt.Printf("%.2f", (currBalance - 0.5 - float64(withdrawlAmt)))
	} else {
		fmt.Printf("%.2f", currBalance)
	}
}
