// Day 1 -
// Part #1 - Find the elf who is carrying the most calorie

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"

	"github.com/adam-lavrik/go-imath/ix"
)

var configFilep1 string = "input.txt"
var elfCalories []int
var calorieCount int = 0

func main() {
	readConfig(configFilep1)
	fmt.Println("D1P1: " + day1P1Answer())
	fmt.Println("D1P2: " + day1P2Answer())
}

func day1P2Answer() string {
	sort.Sort(sort.Reverse(sort.IntSlice(elfCalories)))

	return strconv.Itoa((elfCalories[0] + elfCalories[1] + elfCalories[2]))
}

func day1P1Answer() string {
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
