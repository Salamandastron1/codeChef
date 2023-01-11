package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var ships map[string]string = map[string]string{
	"b": "BattleShip",
	"c": "Cruiser",
	"d": "Destroyer",
	"f": "Frigate",
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	numLines, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))
	for i := 0; i < numLines; i++ {
		scanner.Scan()
		fmt.Println(ships[strings.ToLower(scanner.Text())])
	}
}
