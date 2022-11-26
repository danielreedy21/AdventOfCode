package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	startTime := time.Now()

	// test cases for my bitwise functions
	fmt.Print("Running Main ...\n\n\n")
	fmt.Println("Bitwise function tests: ")
	x := 123
	y := 456
	fmt.Println("x is", x, "\ny is", y)
	fmt.Println("x and y is:", and(x, y))
	fmt.Println("x or y is:", or(x, y))
	fmt.Println("x LSHIFT 2 is:", lshift(x, 2))
	fmt.Println("y RSHIFT 2 is:", rshift(y, 2))
	fmt.Println("not x is:", not(x))
	fmt.Println("not y is:", not(y))

	// answering part 1
	// for part 2, just switch b's signal (line 90) from 14146 to 956

	//reads input
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var lines [339]string
	i := 0
	for scanner.Scan() {
		lines[i] = scanner.Text()
		i++
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// these maps house the calculated wires and the wires that still just hold instructions
	calcWires := map[string]int{}
	instrWires := map[string][]string{}

	// iterate over input lines
	for i := 0; i < len(lines); i++ {
		parsedLine := parseString(lines[i])
		wireName := parsedLine[len(parsedLine)-1]
		// fmt.Println(parsedLine, len(parsedLine))

		// checks if the line on the input is a number while assigning that number to a variable
		// effectively, this grabs the initial assignments
		if num, err := strconv.Atoi(parsedLine[0]); err == nil && len(parsedLine) == 3 {
			calcWires[wireName] = num
		} else {
			instrSlice := parsedLine[0 : len(parsedLine)-2]
			instrWires[wireName] = instrSlice
		}
	}

	// while there are remaining instrucitons, execute the instructions that can be executed
	for len(instrWires) > 0 {
		for key, instruction := range instrWires {

			// case assigning
			if len(instruction) == 1 {
				if num, ok := calcWires[instruction[0]]; ok {
					calcWires[key] = num
					delete(instrWires, key)
				}

				//case NOT
			} else if instruction[0] == "NOT" {
				// if instuction variable exists in calcWires, execute command
				if num, ok := calcWires[instruction[1]]; ok {
					calcWires[key] = not(num)
					delete(instrWires, key)
				}

				// case AND
			} else if instruction[1] == "AND" {
				num1, ok1 := calcWires[instruction[0]]
				num2, ok2 := calcWires[instruction[2]]
				int1, err1 := strconv.Atoi(instruction[0])
				int2, err2 := strconv.Atoi(instruction[2])
				if ok1 && ok2 {
					calcWires[key] = and(num1, num2)
					delete(instrWires, key)
				} else if ok1 && err2 == nil {
					calcWires[key] = and(num1, int2)
					delete(instrWires, key)
				} else if err1 == nil && err2 == nil {
					calcWires[key] = and(int1, int2)
					delete(instrWires, key)
				} else if err1 == nil && ok2 {
					calcWires[key] = and(int1, num2)
					delete(instrWires, key)
				}

				// case OR
			} else if instruction[1] == "OR" {
				num1, ok1 := calcWires[instruction[0]]
				num2, ok2 := calcWires[instruction[2]]
				int1, err1 := strconv.Atoi(instruction[0])
				int2, err2 := strconv.Atoi(instruction[2])
				if ok1 && ok2 {
					calcWires[key] = or(num1, num2)
					delete(instrWires, key)
				} else if ok1 && err2 == nil {
					calcWires[key] = or(num1, int2)
					delete(instrWires, key)
				} else if err1 == nil && err2 == nil {
					calcWires[key] = or(int1, int2)
					delete(instrWires, key)
				} else if err1 == nil && ok2 {
					calcWires[key] = or(int1, num2)
					delete(instrWires, key)
				}

				// case LSHIFT
			} else if instruction[1] == "LSHIFT" {
				if num, ok := calcWires[instruction[0]]; ok {
					shiftNum, _ := strconv.Atoi(instruction[2])
					calcWires[key] = lshift(num, shiftNum)
					delete(instrWires, key)
				}

				// case RSHIFT
			} else if instruction[1] == "RSHIFT" {
				if num, ok := calcWires[instruction[0]]; ok {
					shiftNum, _ := strconv.Atoi(instruction[2])
					calcWires[key] = rshift(num, shiftNum)
					delete(instrWires, key)
				}
			}
		}
	}

	fmt.Printf("\n\nwire A value is: %v", calcWires["a"])

	// calculate execution time
	endTime := time.Now()
	runTime := endTime.Sub(startTime)
	fmt.Println("\n\nExecution Time: ")
	fmt.Println(runTime)
}

// need this function to split input lines into an array of strings
func parseString(line string) []string {
	parsedString := strings.Split(line, " ")
	return parsedString
}

// BITWISE OPERATIONS
func and(num1 int, num2 int) int {
	byte1 := intToByte(num1)
	byte2 := intToByte(num2)
	var newByte [16]int
	for i := 0; i < len(byte1); i++ {
		if byte1[i] == 1 && byte2[i] == 1 {
			newByte[i] = 1
		}
	}
	return byteToInt(newByte)
}

func or(num1 int, num2 int) int {
	byte1 := intToByte(num1)
	byte2 := intToByte(num2)
	var newByte [16]int
	for i := 0; i < len(byte1); i++ {
		if byte1[i] == 1 || byte2[i] == 1 {
			newByte[i] = 1
		}
	}
	return byteToInt(newByte)
}

func lshift(num1 int, shiftNum int) int {
	byte1 := intToByte(num1)
	var newByte [16]int
	//this adds zeros to the nth postitions on the right side of the array
	for i := 1; i <= shiftNum; i++ {
		newByte[len(newByte)-i] = 0
	}
	// iterate over the remaining array elements and update with the shifted value
	for i := 0; i < len(newByte)-shiftNum; i++ {
		newByte[i] = byte1[i+shiftNum]
	}
	return byteToInt(newByte)
}

func rshift(num1 int, shiftNum int) int {
	byte1 := intToByte(num1)
	var newByte [16]int
	//this adds zeros to the nth postitions on the left side of the array
	for i := 0; i < shiftNum; i++ {
		newByte[i] = 0
	}
	// iterate over the remaining array elements and update with the shifted value
	for i := shiftNum; i < len(newByte); i++ {
		newByte[i] = byte1[i-shiftNum]
	}
	return byteToInt(newByte)
}

func not(num1 int) int {
	byte1 := intToByte(num1)
	var newByte [16]int
	for i := 0; i < len(newByte); i++ {
		if byte1[i] == 0 {
			newByte[i] = 1
		}
	}
	return byteToInt(newByte)
}

func intToByte(num int) [16]int {
	var newByte [16]int
	difference := num
	// this is like a while loop
	for difference > 0 {
		for i := 0; i < len(newByte); i++ {
			currPow := 15 - i
			twoPow := int(math.Pow(2, float64(currPow)))
			if difference-twoPow >= 0 {
				newByte[i] = 1
				difference = difference - twoPow
			}
		}
	}
	return newByte
}

func byteToInt(byte1 [16]int) int {
	sum := 0
	for i := 0; i < len(byte1); i++ {
		if byte1[i] == 1 {
			currPow := 15 - i
			twoPow := int(math.Pow(2, float64(currPow)))
			sum = sum + twoPow
		}
	}
	return sum
}
