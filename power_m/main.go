package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	var X, M int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	t := toInt(scanner.Bytes())

	for t > 0 {
		scanner.Scan()
		X = toInt(scanner.Bytes())
		scanner.Scan()
		M = toInt(scanner.Bytes())
		sum := 0
		for X > 0 {
			d := X % 10
			z := M % 4
			if z == 0 && M != 0 {
				z = 4
			}
			d = int(math.Pow(float64(d), float64(z)))
			sum = sum*10 + d%10
			X = X / 10
		}
		if sum%7 == 0 {
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
