package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type coldDayTuple struct {
	first  int
	second int
}

func main() {
	var coldDays coldDayTuple
	s := bufio.NewScanner(os.Stdin)
	s.Scan()
	numLines := parseInt(s.Text())

	for i := 0; i < numLines; i++ {
		s.Scan()
		s.Scan()
		for i, v := range strings.Split(s.Text(), " ") {
			num := parseInt(v)
			if i == 0 {
				coldDays.first = num
				continue
			}
			if coldDays.first > num {
				coldDays.second = coldDays.first
				coldDays.first = num
				continue
			}
			if i == 1 {
				coldDays.second = num
			}

		}
		fmt.Println(coldDays.second)
	}

}

func parseInt(s string) int {
	num, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		panic(err)
	}
	return num
}
