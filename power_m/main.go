package main

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
)

func main() {
	var M int
	var X []byte
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	t := toInt(scanner.Bytes())

	for t > 0 {
		scanner.Scan()
		X = scanner.Bytes()
		scanner.Scan()
		M = toInt(scanner.Bytes())

		for i, v := range X {
			X[i] = byte(int(math.Pow(float64(v-'0'), float64(M)))%10) + '0'
		}
		for i := 0; i < (len(X) / 2); i++ {
			X[i], X[len(X)-i-1] = X[len(X)-i-1], X[i]
		}

		X = bytes.TrimLeft(X, "0")
		fmt.Println(X, toInt(X))
		if toInt(X)%7 == 0 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
		t--
	}
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}
