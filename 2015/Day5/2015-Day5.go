package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	//read input
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//scan over each line in the input
	scanner := bufio.NewScanner(file)
	var lines [1000]string
	i := 0
	for scanner.Scan() {
		lines[i] = scanner.Text()
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//run the checker on each line
	fmt.Println(len(lines))
	niceCounter := 0
	for i := 0; i < len(lines); i++ {
		input := lines[i]
		if naughtOrNice2(input) {
			niceCounter++
		}
	}
	fmt.Println(niceCounter)

	// fmt.Println(naughtOrNice2("qjhvhtzxzqqjkmpb"))
	// fmt.Println(naughtOrNice2("xxyxx"))
	// fmt.Println(naughtOrNice2("uurcxstgmygtbstg"))
	// fmt.Println(naughtOrNice2("ieodomkazucvgmuy"))
}

func naughtOrNice(input string) bool {
	vowelCount := 0
	firstChar := input[0]
	if firstChar == 'a' || firstChar == 'e' || firstChar == 's' || firstChar == 'o' || firstChar == 'u' {
		vowelCount++
	}
	// letterMap := map[byte]int{firstChar: 1}
	repeatCount := 0
	for i := 1; i < len(input); i++ {
		lastChar := input[i-1]
		currChar := input[i]
		// check if currChar and lastChar are repeated
		if lastChar == currChar {
			repeatCount++
		}
		// create string from the last two characters and check that it is not naughty
		last2Char := [2]byte{lastChar, currChar}
		last2str := string(last2Char[:])
		if last2str == "ab" || last2str == "cd" || last2str == "pq" || last2str == "xy" {
			return false
		}
		// calculate the vowel count
		if currChar == 'a' || currChar == 'e' || currChar == 'i' || currChar == 'o' || currChar == 'u' {
			vowelCount++
		}
		// letterMap[currChar]=
		// _, ok := letterMap[currChar]
		// if ok {
		// 	val := letterMap[currChar]
		// 	letterMap[currChar] = val + 1
		// } else {
		// 	letterMap[currChar] = 1
		// }
	}
	// check that there are at least three vowels
	if vowelCount < 3 {
		return false
	}
	// check that a char has repeated at least once
	if repeatCount == 0 {
		return false
	}
	return true
}

func naughtOrNice2(input string) bool {
	fmt.Println(input)
	firstChar := input[0]
	secondChar := input[1]
	last2Char := [2]byte{firstChar, secondChar}
	last2str := string(last2Char[:])
	// create a map that keeps track of the repeat strings
	repeatMap := map[string]int{}
	repeatMap[last2str] = 0

	repeatCount := 0
	duplicateRepeatCount := 0
	for i := 2; i < len(input); i++ {
		lastLastChar := input[i-2]
		lastChar := input[i-1]
		currChar := input[i]
		last2Char = [2]byte{lastChar, currChar}
		last2str = string(last2Char[:])
		// check if last2str exists in the map, else add it
		_, ok := repeatMap[last2str]
		if ok {
			if i-1 >= repeatMap[last2str]+2 {
				duplicateRepeatCount++
			}
		} else {
			repeatMap[last2str] = i - 1
		}
		fmt.Println(repeatMap)

		// check for the seperated repeat
		if currChar == lastLastChar {
			repeatCount++
		}
	}
	fmt.Println(repeatCount)
	if repeatCount == 0 {
		return false
	}
	fmt.Println(duplicateRepeatCount)
	if duplicateRepeatCount == 0 {
		return false
	}
	fmt.Println()
	return true
}


