package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {

	// READING INPUTS INTO ARRAY
	lines := readInput("input.txt")

	numLiteral := 0
	numNotInMemory := 0
	numEncoded := 0

	// iterate over the array
	for _, line := range lines {
		numLiteral = countLiteral(line) + numLiteral
		numNotInMemory = countNotInMemory(line) + numNotInMemory
		numEncoded = lengthEncodedString(line) + numEncoded
	}

	fmt.Println(numLiteral)
	fmt.Println(numNotInMemory)
	fmt.Println(numEncoded)
}

func lengthEncodedString(input string) int {
	// start at 2 because we at least will have two quotes
	count := 2
	for i := 0; i < len(input); i++ {
		currByte := input[i]
		switch currByte {
		case 34: // QUOTATIONS
			count++
		case 92: //BACKSLASHES
			count++
		}
	}
	return count
}

// This function should take in a string and return its length
func countLiteral(input string) int {
	count := 0
	for i := 0; i < len(input); i++ {
		count++
	}
	return count
}

// this function should return the amount of space a character uses in memory
func countNotInMemory(input string) int {
	count := 0
	// iterate over the input string
	for i := 0; i < len(input); i++ {

		currByte := input[i]

		// add to the count depending on what the character is
		switch currByte {
		// CASE QUOTATIONS
		case 34:
			count++
		// CASE BACKSLASH
		case 92:
			nextByte := input[i+1]

			// if the current char is a backslash, then we need to check the next char
			switch nextByte {
			case 120: // HEXCODE
				count += 3
			default: // this will always be a backslash or a quote
				count++
				i++ // I increase the index here because we don't need to check the next character
			}
		}
	}
	return count
}

// This function is my standard for reading in the txt file input by line
func readInput(path string) [300]string {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines [300]string
	i := 0
	for scanner.Scan() {
		lines[i] = scanner.Text()
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return lines
}
