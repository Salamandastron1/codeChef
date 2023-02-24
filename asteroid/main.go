package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type asteroid struct {
	direction int
	size      int
	destroyed bool
	origin    int
}

type finalState struct {
	numAsteroids int
	indices      string
}

func main() {
	// 	Input Format

	// First line will contain T, the number of test cases.
	// Then the test cases follow - each test case has N+1 lines.
	// The first line of input of each test case contains a single integer N.
	// Next NN lines contain two space-separated integers: diri​ and ai​
	// - representing the direction (00 if left, 11 if right)
	// and the size of the asteroid.

	// Output Format

	// For each test case, first output in a single line
	// the number of survived asteroids - xx. In the next line,
	// output xx integers representing the indices of the survived asteroids
	// sorted by increasing ordersorted by increasing order.
	var N int
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	for scanner.Scan() {
		N = toInt(scanner.Bytes())
		xa := make([]asteroid, 0, N)
		// var survivors []asteroid
		var a asteroid
		// the current asteroid must look forward
		// if the current asteroid and the next
		// are moving the same direction the current survives.
		// we can therefore skip ahead.
		for i := 0; i < N; i++ {
			scanner.Scan()
			for j, v := range strings.Split(scanner.Text(), " ") {
				// origin 0 not allowed
				a.origin = i + 1
				if j == 0 {
					a.direction, _ = strconv.Atoi(v)
				} else {
					a.size, _ = strconv.Atoi(v)
				}
			}
			xa = append(xa, a)
			a = asteroid{}
		}
		xa = moveAsteroids(xa)

		var sb strings.Builder
		notDestroyed := finalState{}

		for i := range xa {
			if !xa[i].destroyed {
				notDestroyed.numAsteroids++
				sb.WriteString(fmt.Sprint(xa[i].origin))
				sb.WriteString(" ")
			}
		}
		fmt.Println(notDestroyed.numAsteroids)
		if sb.Len() > 0 {
			fmt.Println(sb.String())
		}

		xa = []asteroid{}
	}
}

func toInt(buf []byte) (n int) {
	for _, v := range buf {
		n = n*10 + int(v-'0')
	}
	return
}

// not accounting for left movement
func moveAsteroids(xa []asteroid) []asteroid {
	var diffDirection bool = true
	for diffDirection {
		var left bool
		var right bool
		i := 0
		for i < len(xa) {
			// if an asteroid is destroyed move on
			if xa[i].destroyed || i == len(xa)-1 {
				i++
				continue
			}
			// keep track of directions of asteroids
			if xa[i].direction == 0 {
				left = true
			} else {
				right = true
			}

			for j := i + 1; j < len(xa); j++ {
				if xa[j].destroyed {
					continue
				}
				if xa[i].direction == xa[j].direction && xa[j].destroyed == false {
					break
				}
				if xa[i].size > xa[j].size {
					xa[j].destroyed = true
					xa[i].size += xa[j].size
					i = -1
					break
				} else if xa[i].size == xa[j].size {
					xa[j].destroyed = true
					xa[i].destroyed = true
					i = -1
					break
				} else if xa[i].size < xa[j].size {
					xa[i].destroyed = true
					xa[j].size += xa[i].size
					i = -1
					break
				}
			}
			i++
		}
		// check to see if asteroids are traveling
		// in same direction
		// if they are, end the loop
		if left != right {
			diffDirection = false
		}
	}
	return xa
}
