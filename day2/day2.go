// Day 2 -
// Part #1 -
// Part #2 -

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var configFilep1 string = "input.txt"
var ourMoves []string
var opponentMoves []string

func main() {
	readConfig(configFilep1)

	fmt.Println(ourMoves)

	fmt.Println("D1P1: " + day2P1Answer())
	fmt.Println("D1P2: " + day2P2Answer())
}

func day2P2Answer() string {
	var sumMyMoves int = 0
	for i, v := range ourMoves {
		switch v {
		case "X":
			if opponentMoves[i] == "A" {
				sumMyMoves += 3
			} else if opponentMoves[i] == "B" {
				sumMyMoves += 1
			} else {
				sumMyMoves += 2
			}
		case "Y":
			if opponentMoves[i] == "A" {
				sumMyMoves += 4
			} else if opponentMoves[i] == "B" {
				sumMyMoves += 5
			} else {
				sumMyMoves += 6
			}
		case "Z":
			if opponentMoves[i] == "A" {
				sumMyMoves += 8
			} else if opponentMoves[i] == "B" {
				sumMyMoves += 9
			} else {
				sumMyMoves += 7
			}
		}
	}
	return strconv.Itoa(sumMyMoves)
}

func day2P1Answer() string {

	var sumMyMoves int = 0

	// Sum of what I throw
	for i, v := range ourMoves {
		fmt.Println(v)
		switch v {
		case "X":
			sumMyMoves += 1
		case "Y":
			sumMyMoves += 2
		case "Z":
			sumMyMoves += 3
		}
		sumMyMoves += calcWinner(i)
	}

	return strconv.Itoa(sumMyMoves)
}

func calcWinner(i int) int {
	switch ourMoves[i] {
	case "X":
		switch opponentMoves[i] {
		case "A":
			return 3
		case "B":
			return 0
		case "C":
			return 6
		}
	case "Y":
		switch opponentMoves[i] {
		case "A":
			return 6
		case "B":
			return 3
		case "C":
			return 0
		}
	case "Z":
		switch opponentMoves[i] {
		case "A":
			return 0
		case "B":
			return 6
		case "C":
			return 3
		}

	}
	return 1
}

// Read in configuration
func readConfig(f string) {
	readFile, err := os.Open(f)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		s := strings.Fields(fileScanner.Text())
		opponentMoves = append(opponentMoves, s[0])
		ourMoves = append(ourMoves, s[1])
	}
	readFile.Close()
}
