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
		// Ai,j​=1, for all 1≤i≤N and 1≤j≤M,
		// except for i=N and j=M.
		// AN,M=0AN,M​=0.
		scanner.Scan()
		n = toInt(scanner.Bytes())
		scanner.Scan()
		m = toInt(scanner.Bytes())
		fmt.Println((m * (m + 1) * (n) * (n + 1) / 4) - (m * n))
		t--
		// (𝑚+1)𝑚/2⋅(𝑛+1)𝑛/2
		// https://en.wikipedia.org/wiki/Binomial_distribution
	}
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}
