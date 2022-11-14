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
	//read input 
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	//scan over each line in the input
	scanner := bufio.NewScanner(file)
	var lines [300]string
	i := 0
	for scanner.Scan() {
		lines[i] = scanner.Text()
		i++;
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// create lights 2D array
	var lights [1000][1000]int
	//iterate through each line and grab the parameters from it
	for i:=0;i<len(lines);i++ {
		input := lines[i]
		inputList := strings.Split(input, " ")
		firstCoord := inputList[len(inputList)-3]
		secondCoord := inputList[len(inputList)-1]
		firstCoordList := strings.Split(firstCoord, ",")
		secondCoordList := strings.Split(secondCoord, ",")
		x1, _ := strconv.Atoi(firstCoordList[0])
		y1, _ := strconv.Atoi(firstCoordList[1])
		x2, _ := strconv.Atoi(secondCoordList[0])
		y2, _ := strconv.Atoi(secondCoordList[1])
		command := inputList[len(inputList)-4]
		// once I have grabbed all the parameters, call the correct function with them
		if command == "on" {
			lights = turnOn(lights,x1,y1,x2,y2)
		} else if command == "off" {
			lights = turnOff(lights,x1,y1,x2,y2)
		} else if command == "toggle" {
			lights = toggle(lights,x1,y1,x2,y2)
		}
	}
	fmt.Println(countOn(lights))
}





func countOn(lights [1000][1000]int) int {
	count := 0
	for i:=0; i<len(lights); i++ {
		for j:=0; j<len(lights[0]); j++ {
			count = count + lights[i][j]
		}
	}
	return count
}
func turnOn(lights [1000][1000]int, x1 int, y1 int, x2 int, y2 int) [1000][1000]int {
	if x1<=x2 && y1<=y2 {
		for i:=x1; i<=x2; i++ {
			for j:=y1; j<=y2; j++ {
				lights[i][j]= lights[i][j] + 1
			}
		}
		return lights
	} else if x1>x2 && y1<=y2 {
		for i:=x1; i>=x2; i-- {
			for j:=y1; j<=y2; j++ {
				lights[i][j]= lights[i][j] + 1
			}
		}
		return lights
	} else if x1>x2 && y1>y2 {
		for i:=x1; i>=x2; i-- {
			for j:=y1; j>=y2; j-- {
				lights[i][j]= lights[i][j] + 1
			}
		}
		return lights
	} else if x1<=x2 && y1>y2 {
		for i:=x2; i<=x1; i++ {
			for j:=y1; j>=y2; j-- {
				lights[i][j]= lights[i][j] + 1
			}
		}
		return lights
	}
	return lights
	// smallX :=0
	// largeX :=0
	// smallY :=0
	// largeY :=0
	// xDirection := 1
	// yDirection := -1

	// if x1>x2 {
	// 	smallX = x2
	// 	largeX = x1
	// } else {
	// 	smallX = x1
	// 	largeX = x2
	// }
	// if y1>y2 {
	// 	smallY = y2
	// 	largeY = y1
	// } else {
	// 	smallY = y1
	// 	largeY = y2
	// }
	// for i:=smallX; i<=largeX; i++ {
	// 	for j:=smallY; j<=largeY; j++ {
	// 		lights[i][j]=1
	// 	}
	// }
	// return lights
}
func turnOff(lights [1000][1000]int, x1 int, y1 int, x2 int, y2 int) [1000][1000]int {
	if x1<=x2 && y1<=y2 {
		for i:=x1; i<=x2; i++ {
			for j:=y1; j<=y2; j++ {
				if lights[i][j]>0 {
					lights[i][j] = lights[i][j] - 1
				}
			}
		}
		return lights
	} else if x1>x2 && y1<=y2 {
		for i:=x1; i>=x2; i-- {
			for j:=y1; j<=y2; j++ {
				if lights[i][j]>0 {
					lights[i][j] = lights[i][j] - 1
				}
			}
		}
		return lights
	} else if x1>x2 && y1>y2 {
		for i:=x1; i>=x2; i-- {
			for j:=y1; j>=y2; j-- {
				if lights[i][j]>0 {
					lights[i][j] = lights[i][j] - 1
				}
			}
		}
		return lights
	} else if x1<=x2 && y1>y2 {
		for i:=x2; i<=x1; i++ {
			for j:=y1; j>=y2; j-- {
				if lights[i][j]>0 {
					lights[i][j] = lights[i][j] - 1
				}
			}
		}
		return lights
	}
	return lights
}
func toggle(lights [1000][1000]int, x1 int, y1 int, x2 int, y2 int) [1000][1000]int {
	if x1<=x2 && y1<=y2 {
		for i:=x1; i<=x2; i++ {
			for j:=y1; j<=y2; j++ {
				lights[i][j] = lights[i][j] + 2
			}
		}
		return lights
	} else if x1>x2 && y1<=y2 {
		for i:=x1; i>=x2; i-- {
			for j:=y1; j<=y2; j++ {
				lights[i][j] = lights[i][j] + 2
			}
		}
		return lights
	} else if x1>x2 && y1>y2 {
		for i:=x1; i>=x2; i-- {
			for j:=y1; j>=y2; j-- {
				lights[i][j] = lights[i][j] + 2
			}
		}
		return lights
	} else if x1<=x2 && y1>y2 {
		for i:=x2; i<=x1; i++ {
			for j:=y1; j>=y2; j-- {
				lights[i][j] = lights[i][j] + 2
			}
		}
		return lights
	}
	return lights
}


