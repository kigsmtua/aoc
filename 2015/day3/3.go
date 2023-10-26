package main

import (
	"fmt"
	"os"
)

var Directions map[rune]complex128

func init() {
	Directions = map[rune]complex128{
		'^': 0 + 1i,
		'>': 1 + 0i,
		'v': 0 - 1i,
		'<': -1 + 0i,
	}
}

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Heyooo, that file is not found")
	}
	movements := string(data)
	fmt.Println(problem1(movements))
	fmt.Println(problem2(movements))
}

func problem1(movements string) int {
	visited := make(map[complex128]bool)
	position := complex(0, 0)
	visited[position] = true

	for _, char := range movements {
		position += Directions[char]
		visited[position] = true
	}
	return len(visited)
}

func problem2(movements string) int {
	visited := make(map[complex128]bool)
	santa := complex(0, 0)
	santaBot := complex(0, 0)
	visited[santa] = true
	visited[santaBot] = true

	for pos, char := range movements {
		if pos%2 == 0 {
			santa += Directions[char]
			visited[santa] = true
		} else {
			santaBot += Directions[char]
			visited[santaBot] = true
		}
	}
	return len(visited)
}
