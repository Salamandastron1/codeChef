package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
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
		var n string
		_, err = scanf("%s\n", &n)
		if err != nil {
			fmt.Println(err)
		}
		res := ""
		for i := len(n) - 1; i >= 0; i-- {
			res += string(n[i])
		}
		num, _ := strconv.Atoi(res)
		_, err = printf("%d\n", num)
		if err != nil {
			fmt.Println(err)
		}
	}
}
