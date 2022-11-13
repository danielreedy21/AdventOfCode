package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
)

func main() {

	// read input file 
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	
	// scan over file and add lines to an array
	scanner := bufio.NewScanner(file)
	var lines [1000]string
	i := 0
	for scanner.Scan() {
		lines[i] = scanner.Text()
		i++;
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// print the calculation
	fmt.Println(getTotal(lines))
}

func wrappingPaper(l int, w int, h int) int {
	side1 := l*w
	side2 := w*h
	side3 := h*l
	SA := (side1+side2+side3)*2
	slack := 0
	if side1 <= side2 && side1 <= side3 {
		slack = side1
	} else if side2 <= side1 && side2 <= side3 {
		slack = side2
	} else {
		slack = side3
	}
	return SA + slack
}

func ribbons(l int, w int, h int) int {
	total := 0
	cubic := l*w*h
	perim1 := 2*(l+w)
	perim2 := 2*(l+h)
	perim3 := 2*(w+h)
	if perim1 <= perim2 && perim1 <= perim3 {
		total = total + perim1
	} else if perim2 <= perim1 && perim2 <= perim3 {
		total = total + perim2
	} else {
		total = total + perim3
	}
	total = total + cubic
	return total
}

func getTotal(lines [1000]string) [2]int {
	total := 0
	ribTotal := 0
	for i := 0; i<len(lines); i++ {
		line := lines[i]
		values := strings.Split(line, "x")
		var valInts [3]int
		for j := 0; j<len(values); j++ {
			intVal, err := strconv.Atoi(values[j]);  
			if err != nil {
				panic(err)
			}
			valInts[j] = intVal
		}
		l := valInts[0]
		w := valInts[1]
		h := valInts[2]
		total = total + wrappingPaper(l,w,h)
		ribTotal = ribTotal + ribbons(l,w,h)
	}
	return [2]int{total,ribTotal}
}