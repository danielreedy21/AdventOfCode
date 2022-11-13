package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
)

func main() {
	//read input file
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// scan over input file and turn it into a string
	scanner := bufio.NewScanner(file)
	var input string
	for scanner.Scan(){
		input = scanner.Text()
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// run solution
	fmt.Println(numDuplicateHouses(input))
}

func numDuplicateHouses(input string) int {
	santaCoords := [2]int{0,0}
	roboCoords := [2]int{0,0}
	houseSet := map[[2]int]bool{santaCoords: true}
	// how to add to the set
	// houseSet[[2]int{1,1}]=true
	// how to check if an element exists in the set
	// _,ok := houseSet[[2]int{0,0}]

	for i:=0; i<len(input); i++{
		currChar := input[i]
		if i%2 == 0 {
			if currChar == '>' {
				santaCoords[0] = santaCoords[0] + 1
			} else if currChar == '<' {
				santaCoords[0] = santaCoords[0] - 1
			} else if currChar == '^' {
				santaCoords[1] = santaCoords[1] + 1
			} else if currChar == 'v' {
				santaCoords[1] = santaCoords[1] - 1
			} else {
				fmt.Println("Invalid Input")
			}
			houseSet[santaCoords]=true
		} else {
			if currChar == '>' {
				roboCoords[0] = roboCoords[0] + 1
			} else if currChar == '<' {
				roboCoords[0] = roboCoords[0] - 1
			} else if currChar == '^' {
				roboCoords[1] = roboCoords[1] + 1
			} else if currChar == 'v' {
				roboCoords[1] = roboCoords[1] - 1
			} else {
				fmt.Println("Invalid Input")
			}
			houseSet[roboCoords]=true
		}
	}
	return len(houseSet)
}