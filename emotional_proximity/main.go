package main

import (
	"bufio"
	"fmt"
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
	var p, x, y, z float64
	scanf("%d\n", &T)

	for T > 0 {
		scanf("%f %f %f %f\n", &p, &x, &y, &z)
		if z == 0 {
			printf("%.10f\n", p*(1-(x/100)))
		} else {
			printf("%.10f\n", p*(1+(y/100)))
		}
		T--
	}
}
