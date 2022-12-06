package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var configFilep1 string = "input.txt"
var elfList []elves

type elves struct {
	fRange string
	sRange string
}

func main() {
	readConfig(configFilep1)

	fmt.Println("D3P1: " + day3P1Answer())
	fmt.Println("D3P2: " + day3P2Answer())
}

func day3P2Answer() string {
	var count int = 0

	for _, v := range elfList {
		fSplit := strings.Split(v.fRange, "-")
		fL, _ := strconv.Atoi(fSplit[0])
		fR, _ := strconv.Atoi(fSplit[1])
		sSplit := strings.Split(v.sRange, "-")
		sL, _ := strconv.Atoi(sSplit[0])
		sR, _ := strconv.Atoi(sSplit[1])

		if (fR >= sL && fR <= sR) || (fL >= sL && fL <= sR) || (sL >= fL && sL <= fR) {
			count += 1
		}
	}
	return strconv.Itoa(count)
}

func day3P1Answer() string {
	var count int = 0
	for _, v := range elfList {
		fSplit := strings.Split(v.fRange, "-")
		fL, _ := strconv.Atoi(fSplit[0])
		fR, _ := strconv.Atoi(fSplit[1])
		sSplit := strings.Split(v.sRange, "-")
		sL, _ := strconv.Atoi(sSplit[0])
		sR, _ := strconv.Atoi(sSplit[1])

		if (fL <= sL && fR >= sR) || (sL <= fL && sR >= fR) {
			count += 1
		}
	}
	return strconv.Itoa(count)
}

// Read in configuration
func readConfig(f string) {
	readFile, err := os.Open(f)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		s := strings.Split(fileScanner.Text(), ",")
		elfList = append(elfList, elves{s[0], s[1]})
	}
	readFile.Close()
}
