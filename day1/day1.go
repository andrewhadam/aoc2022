// Day 1 -
// Part #1 - Find the elf who is carrying the most calorie

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/adam-lavrik/go-imath/ix"
)

var configFilep1 string = "input.txt"
var elfCalories []int
var calorieCount int = 0

func main() {
	readConfig(configFilep1)
	fmt.Println("D1P1: " + day1Answer())
}

func day1Answer() string {
	return strconv.Itoa(ix.MaxSlice(elfCalories))
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
		if fileScanner.Text() == "" {
			elfCalories = append(elfCalories, calorieCount)
			calorieCount = 0
			continue
		} else {
			calVal, err := strconv.Atoi(fileScanner.Text())
			if err != nil {
				fmt.Println("Unable to interpret string")
			}
			calorieCount += calVal
		}
	}
	readFile.Close()
}
