package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

type rucksack struct {
	firstCompartment  string
	secondCompartment string
	common            string
}

type threeWayRucksack struct {
	firstCompartment  string
	secondCompartment string
	thirdCompartment  string
	common            string
}

var configFilep1 string = "input.txt"
var myRucksacks []rucksack
var myThreeWayRucksack []threeWayRucksack

func main() {
	readConfig(configFilep1)
	readThreeWayConfig(configFilep1)

	//fmt.Println("D3P1: " + day3P1Answer())
	fmt.Println("D3P2: " + day3P2Answer())
}

func day3P2Answer() string {
	var sumThreeSacks int = 0
	for _, v := range myThreeWayRucksack {
		r := []rune(v.common)
		var r3 = r[0]
		fmt.Println("Day2 answer " + string(r3))
		fmt.Println(r3)
		if unicode.IsUpper(r3) {
			// This is -64 to set to 1 and then + 26 to set to 27 as a baseline
			sumThreeSacks += (int(r3) - 38)
			fmt.Println("It is uppercase")
		} else {
			sumThreeSacks += (int(r3) - 96)
			fmt.Println("It is not uppercase")
			fmt.Println(int(r3) - 96)

		}
	}
	return strconv.Itoa(sumThreeSacks)
}

func day3P1Answer() string {
	var sumSacks int = 0
	for _, v := range myRucksacks {
		r := []rune(v.common)
		var r1 = r[0]
		if unicode.IsUpper(r1) {
			// This is -64 to set to 1 and then + 26 to set to 27 as a baseline
			sumSacks += (int(r1) - 38)

		} else {
			sumSacks += (int(r1) - 96)
		}
	}
	return strconv.Itoa(sumSacks)
}

// Read in configuration
func readConfig(f string) {
	readFile, err := os.Open(f)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		fc := fileScanner.Text()[:(len(fileScanner.Text()) / 2)]
		sc := fileScanner.Text()[(len(fileScanner.Text()) / 2):]
		myRucksacks = append(myRucksacks, rucksack{fc, sc, compare(fc, sc)})
	}
	readFile.Close()
}

func readThreeWayConfig(f string) {
	readFile, err := os.Open(f)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewReader(readFile)

	for {
		fc, _, err := fileScanner.ReadLine()
		if err != nil || len(fc) == 0 {
			break
		}
		sc, _, _ := fileScanner.ReadLine()
		tc, _, _ := fileScanner.ReadLine()

		fmt.Printf("FC: %s \nsc: %s\ntc: %s\n", fc, sc, tc)

		myThreeWayRucksack = append(myThreeWayRucksack, threeWayRucksack{string(fc), string(sc), string(tc), compareThree(string(fc), string(sc), string(tc))})

	}
	readFile.Close()
}

func compare(f, s string) string {
	//	fmt.Println("Comparing " + f + " and " + s)
	for _, fs := range f {
		for _, ss := range s {
			if string(fs) == string(ss) {
				return string(fs)
			}
		}
	}
	return "impossibru!"
}

func compareThree(f, s, t string) string {
	fmt.Println("Comparing " + f + " and " + s + " and " + t)
	for _, fs := range f {
		for _, ss := range s {
			for _, ts := range t {
				if string(fs) == string(ss) && string(fs) == string(ts) {
					fmt.Println(string(fs))
					return string(fs)
				}
			}
		}
	}
	return "impossibru!"
}
