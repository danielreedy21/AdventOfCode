package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	startTime := time.Now()
	x := [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 0, 1, 1}
	y := [16]int{0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 1, 0, 0, 0}
	fmt.Println("Running Main ...")
	fmt.Println("x is :")
	fmt.Println(x)
	fmt.Println("y is :")
	fmt.Println(y)
	fmt.Println("x and y is:")
	fmt.Println(and(x, y))
	fmt.Println("x or y is:")
	fmt.Println(or(x, y))
	fmt.Println("16 byte of 234 is :")
	fmt.Println(intToByte(234))
	fmt.Println("x as an int is:")
	fmt.Println(byteToInt(x))

	endTime := time.Now()
	runTime := endTime.Sub(startTime)
	fmt.Println("\n\n\nExecution Time: ")
	fmt.Println(runTime)
}

func and(byte1 [16]int, byte2 [16]int) [16]int {
	var newByte [16]int
	for i := 0; i < len(byte1); i++ {
		if byte1[i] == 1 && byte2[i] == 1 {
			newByte[i] = 1
		}
	}
	return newByte
}

func or(byte1 [16]int, byte2 [16]int) [16]int {
	var newByte [16]int
	for i := 0; i < len(byte1); i++ {
		if byte1[i] == 1 || byte2[i] == 1 {
			newByte[i] = 1
		}
	}
	return newByte
}

func lshift(byte1 [16]int, shiftNum int) [16]int {
	var newByte [16]int
	//this adds zeros to the nth postitions on the right side of the array
	for i := 1; i <= shiftNum; i++ {
		newByte[len(newByte)-i] = 0
	}
	// iterate over the remaining array elements and update with the shifted value
	for i := 0; i < len(newByte)-shiftNum; i++ {
		newByte[i] = byte1[i+shiftNum]
	}
	return newByte
}

func rshift(byte1 [16]int, shiftNum int) [16]int {
	var newByte [16]int
	//this adds zeros to the nth postitions on the left side of the array
	for i := 0; i < shiftNum; i++ {
		newByte[i] = 0
	}
	// iterate over the remaining array elements and update with the shifted value
	for i := shiftNum; i < len(newByte); i++ {
		newByte[i] = byte1[i-shiftNum]
	}
	return newByte
}

func not(byte1 [16]int) [16]int {
	var newByte [16]int
	for i := 0; i < len(newByte); i++ {
		if byte1[i] == 0 {
			newByte[i] = 1
		}
	}
	return newByte
}

func byteToNum(byte1 [16]int) int {
	sum := 0
	for i := 0; i < len(byte1); i++ {
		if byte1[i] == 1 {
			sum = sum + int(math.Pow(2, float64(15-i)))
		}
	}
	return sum
}

func intToByte(num int) [16]int {
	var newByte [16]int
	difference := num
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
