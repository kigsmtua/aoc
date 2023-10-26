package main

import (
	"fmt"
	"os"
)

func main() {
	data, err := os.ReadFile("input.txt")

	if err != nil {
		fmt.Println("Heyooo, that file is not found")
	}

	movements := string(data)
	visited := make(map[complex128]bool)
	position := complex(0, 0)
	visited[position] = true

	directions := map[rune]complex128{
		'^': 0 + 1i,
		'>': 1 + 0i,
		'v': 0 - 1i,
		'<': -1 + 0i,
	}

	for _, char := range movements {
		position += directions[char]
		visited[position] = true
	}

	fmt.Println(visited)
	fmt.Println(len(visited))
}
