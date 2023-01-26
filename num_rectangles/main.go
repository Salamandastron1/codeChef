package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n, m int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	t := toInt(scanner.Bytes())

	for t > 0 {
		scanner.Scan()
		n = toInt(scanner.Bytes())
		scanner.Scan()
		m = toInt(scanner.Bytes())
		fmt.Println((m * (m + 1) * (n) * (n + 1) / 4) - (m * n))
		t--
	}
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}
