package main

import (
	"fmt"
	"os"
)

func main() {
	//Lets read the entire file into memory anyway
	data, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("unable to open file, please specify correct file location")
	}
	fmt.Printf("The answer for question 1 is %d \n", question1(string(data)))
	fmt.Printf("The answer for question 2 is %d 	\n", question2(string(data)))

}

// question 1
func question1(input string) int {
	currentFloor := 0
	for _, char := range input {
		if char == '(' {
			currentFloor = currentFloor + 1
		} else {
			currentFloor = currentFloor - 1
		}
	}
	return currentFloor
}

// question 2
func question2(input string) int {
	currentFloor := 0
	baseMentPosition := 0
	for pos, char := range input {
		if char == '(' {
			currentFloor = currentFloor + 1
		} else {
			currentFloor = currentFloor - 1
		}
		if currentFloor == -1 {
			baseMentPosition = pos + 1
			break
		}
	}
	return baseMentPosition
}
