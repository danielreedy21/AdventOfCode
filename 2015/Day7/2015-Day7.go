package main

import (
	"fmt"
	"math"
	"time"
)

func main() {
	startTime := time.Now()

	// actual executables
	fmt.Print("Running Main ...\n\n\n")
	x := 123
	y := 456
	fmt.Println("x is", x, "\ny is", y)
	fmt.Println("x and y is:", and(x, y))
	fmt.Println("x or y is:", or(x, y))
	fmt.Println("x LSHIFT 2 is:", lshift(x, 2))
	fmt.Println("y RSHIFT 2 is:", rshift(y, 2))
	fmt.Println("not x is:", not(x))
	fmt.Println("not y is:", not(y))

	// calculate execution time
	endTime := time.Now()
	runTime := endTime.Sub(startTime)
	fmt.Println("\n\nExecution Time: ")
	fmt.Println(runTime)
}

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
