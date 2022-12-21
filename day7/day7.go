package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var configFilep1 string = "input.txt"

// var currWorkDir string
var currWorkDirRef *Folder
var myOS []Folder

type File struct {
	name string
	size int64
}

type Folder struct {
	name      string
	files     []File
	folders   []Folder
	parentDir *Folder
	index     int
	size      int64
}

//type Folders []Folder

func main() {

	readConfig(configFilep1)
	fmt.Println("D5P1: " + day7P1Answer())

	fmt.Println("D5P2: " + day7P2Answer())
}

func day7P2Answer() string {
	var answer string = ""

	return answer
}

func day7P1Answer() string {
	var answer string = ""
	fmt.Println()
	fmt.Println()
	fmt.Println()

	fmt.Println(myOS[0].files)

	return answer
}

func readConfig(f string) {
	readFile, err := os.Open(f)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)

	for fileScanner.Scan() {
		fmt.Println("EXECUTING THE FOLLOWING:")
		fmt.Println(fileScanner.Text())
		fmt.Println("CURRENT WORKING DIRECTORY")
		fmt.Println(currWorkDirRef)
		switch {
		case strings.Contains(fileScanner.Text(), "$ cd /"):
			var f Folder = Folder{"/", nil, nil, nil, 0, 0}

			//			fmt.Println(currWorkDirRef)
			//"/", nil, nil, nil, 0)
			myOS = append(myOS, f)
			currWorkDirRef = &myOS[0]
			fmt.Println(len(myOS))
			fmt.Println("CHecking size")
			//			var refTest Folder = myOS[0]
			//			var refTest2 *Folder = &refTest
			//			fmt.Println(refTest)
			//			fmt.Println(refTest2)
		case strings.Contains(fileScanner.Text(), "cd .."):
			currWorkDirRef = currWorkDirRef.parentDir
		// Assume that everything else is changing to a directory
		case strings.Contains(fileScanner.Text(), "cd "):
			changeDirectory(fileScanner.Text())
		case strings.Contains(fileScanner.Text(), "$ ls"):
			continue
		default:
			parseLS(fileScanner.Text())
		}
	}
	readFile.Close()
}

func changeDirectory(s string) {
	var t = strings.Split(s, " ")
	fmt.Println(t[2])

	for _, v := range currWorkDirRef.folders {
		fmt.Println(v.name)
		if v.name == t[2] {
			fmt.Println("Found it")
			fmt.Println(currWorkDirRef)

			currWorkDirRef = &v
			break
			fmt.Println(currWorkDirRef)
		}
		fmt.Println(currWorkDirRef)

	}
	fmt.Println("Changing Directory")
	fmt.Println(currWorkDirRef)
}

func parseLS(s string) {
	// make a directory
	var t = strings.Split(s, " ")
	fmt.Println("PRINTING T IN PARSELS")
	fmt.Println(t)
	fmt.Println(t[1])
	//	fmt.Println(myOS[currWorkDirRef.index].folders)
	fmt.Println(currWorkDirRef.index)

	if strings.Contains(s, "dir ") {
		var u = Folder{t[1], nil, nil, currWorkDirRef, len(currWorkDirRef.folders), 0}
		//currWorkDirRef = append(myOS[currWorkDirRef.index].folders, u)
		myOS[currWorkDirRef.index].folders = append(myOS[currWorkDirRef.index].folders, u)
	} else {
		fmt.Println("Trying to create a file")
		fmt.Println(currWorkDirRef.files)
		fmt.Println(currWorkDirRef.size)
		//currWorkDirRef = &myOS[0]
		fmt.Println(currWorkDirRef.index)

		// We need to create some files
		n, _ := strconv.ParseInt(t[0], 10, 32)
		var f File = File{t[1], n}
		myOS[currWorkDirRef.index].files = append(myOS[currWorkDirRef.index].files, f)
		//currWorkDirRef.files = append(&currWorkDirRef.files, f)
		myOS[currWorkDirRef.index].size += n
		//		currWorkDirRef.size =currWorkDirRef.size + n
	}
}
