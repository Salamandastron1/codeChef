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
	var a, b, c, d, e, f, g int
	scanf("%d\n", &T)

	for T > 0 {
		scanf("%d %d %d %d %d %d %d\n", &a, &b, &c, &d, &e, &f, &g)
		printf("%d\n", ((a/2)+(a%2))*(b+c+d+e+f+g))
		T--
	}
}
