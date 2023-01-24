package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type fn func([]string) bool

var funcMap map[string]fn = map[string]fn{
	"part0": part0,
	"part1": part1,
	"part2": part2,
	"part3": part3,
	"part4": part4,
	"part5": part5,
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for scanner.Scan() {
		if canReachZero(strings.Split(strings.TrimSpace(scanner.Text()), " ")) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func canReachZero(xs []string) bool {
	count := 6
	for i := 0; i < count; i++ {
		if funcMap[fmt.Sprintf("part%v", i)](xs) {
			return true
		}
	}
	return false
}

func part0(xs []string) bool {
	aggregate := 0
	for _, v := range xs {
		num, _ := strconv.Atoi(v)
		aggregate += num
	}

	return aggregate == 0
}

func part1(xs []string) bool {
	aggregate := 0
	for _, v := range xs {
		num, _ := strconv.Atoi(v)
		aggregate -= num
	}

	return aggregate == 0
}
func part2(xs []string) bool {
	aggregate := 0
	xi := []int{}
	for _, v := range xs {
		num, _ := strconv.Atoi(v)
		xi = append(xi, num)
	}
	aggregate = xi[0] - xi[1] + xi[2]
	return aggregate == 0
}

func part3(xs []string) bool {
	aggregate := 0
	xi := []int{}
	for _, v := range xs {
		num, _ := strconv.Atoi(v)
		xi = append(xi, num)
	}
	aggregate = xi[0] + xi[1] - xi[2]
	return aggregate == 0
}

func part4(xs []string) bool {
	aggregate := 0
	xi := []int{}
	for _, v := range xs {
		num, _ := strconv.Atoi(v)
		xi = append(xi, num)
	}
	aggregate = xi[1] + xi[2] - xi[0]
	return aggregate == 0
}

func part5(xs []string) bool {
	aggregate := 0
	xi := []int{}
	for _, v := range xs {
		num, _ := strconv.Atoi(v)
		xi = append(xi, num)
	}
	aggregate = xi[0] - xi[1] - xi[2]
	return aggregate == 0
}
