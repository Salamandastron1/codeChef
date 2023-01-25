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
	var T, n, x, c int
	scanf("%d\n", &T)
	for T > 0 {
		scanf("%d %d\n", &n, &x)
		days := make([]int, n, n)
		for i := 0; i < n; i++ {
			scanf("%d\n", &days[i])
			if days[i] > 0 {
				c += x
				continue
			}
			if i > 0 && days[i-1] == 1 {
				c += x
			}
		}
		printf("%d\n", c)
		c = 0
		T--
	}
}
