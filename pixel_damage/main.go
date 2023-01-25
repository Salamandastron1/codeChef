package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

var reader *bufio.Reader = bufio.NewReader(os.Stdin)
var writer *bufio.Writer = bufio.NewWriter(os.Stdout)

func printf(f string, a ...interface{}) (n int, err error) {
	return fmt.Fprintf(writer, f, a...)
}
func scanf(f string, a ...interface{}) (n int, err error) {
	return fmt.Fscanf(reader, f, a...)
}

func main() {
	defer writer.Flush()
	var T int
	var h, w, x, y, k int
	scanf("%d\n", &T)

	for 0 < T {
		//  HH, WW, XX, YY and KK.
		scanf("%d %d %d %d %d\n", &h, &w, &x, &y, &k)

		if math.Sqrt(math.Pow(float64(h-y), 2)+math.Pow(float64(x-w), 2)) < float64(k) {
			printf("%d\n", 1)
		} else {
			printf("%d\n", 0)
		}
		T--
	}
}
