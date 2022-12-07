package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

var configFilep1 string = "input.txt"

var stacks [9]string

func main() {
	initializeStacks()
	readConfig(configFilep1)
	fmt.Println("D5P1: " + day5P1Answer())

	initializeStacks()
	readConfigP2(configFilep1)
	fmt.Println("D5P2: " + day5P2Answer())
}

func initializeStacks() {
	stacks[0] = "QWPSZRHD"
	stacks[1] = "VBRWQHF"
	stacks[2] = "CVSH"
	stacks[3] = "HFG"
	stacks[4] = "PGJBZ"
	stacks[5] = "QTJHWFL"
	stacks[6] = "ZTWDLVJN"
	stacks[7] = "DTZCJGHF"
	stacks[8] = "WPVMBH"
}

func day5P2Answer() string {
	var answer string = ""
	for _, v := range stacks {
		answer += string(v[len(v)-1])
	}
	return answer
}

func day5P1Answer() string {
	var answer string = ""
	for _, v := range stacks {
		answer += string(v[len(v)-1])
	}
	return answer
}

// Read in configuration
func readConfig(f string) {
	readFile, err := os.Open(f)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		parseIt := regexp.MustCompile(`move\s(\d+)\sfrom\s(\d+)\sto\s(\d+)`)
		res := parseIt.FindAllStringSubmatch(fileScanner.Text(), -1)
		numToMove, _ := strconv.Atoi(res[0][1])
		fromStack, _ := strconv.Atoi(res[0][2])
		targetStack, _ := strconv.Atoi(res[0][3])

		for i := 0; i < numToMove; i++ {
			firstChar := string(stacks[fromStack-1][len(stacks[fromStack-1])-1])
			stacks[fromStack-1] = stacks[fromStack-1][:len(stacks[fromStack-1])-1]
			stacks[targetStack-1] = stacks[targetStack-1] + firstChar
		}
	}
	readFile.Close()
}

// Read in configuration
func readConfigP2(f string) {
	fmt.Println(stacks)
	readFile, err := os.Open(f)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		parseIt := regexp.MustCompile(`move\s(\d+)\sfrom\s(\d+)\sto\s(\d+)`)
		res := parseIt.FindAllStringSubmatch(fileScanner.Text(), -1)
		numToMove, _ := strconv.Atoi(res[0][1])
		fromStack, _ := strconv.Atoi(res[0][2])
		targetStack, _ := strconv.Atoi(res[0][3])

		blockString := string(stacks[fromStack-1][len(stacks[fromStack-1])-numToMove:])
		stacks[fromStack-1] = stacks[fromStack-1][:len(stacks[fromStack-1])-numToMove]
		stacks[targetStack-1] = stacks[targetStack-1] + blockString
	}
	readFile.Close()
}
