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
	_, err := scanf("%d\n", &T)
	if err != nil {
		fmt.Println(err)
	}
	for i := 0; i < T; i++ {
		var n, k int
		_, err = scanf("%d %d\n", &n, &k)
		if err != nil {
			fmt.Println(err)
		}
		_, err = printf("%d\n", n+k)
		if err != nil {
			fmt.Println(err)
		}
	}
}
