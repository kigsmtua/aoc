package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	data, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Please input the correct path to the file")
	}
	fileScanner := bufio.NewScanner(data)
	fileScanner.Split(bufio.ScanLines)
	var fileContents []string
	for fileScanner.Scan() {
		fileContents = append(fileContents, fileScanner.Text())
	}

	totalGiftWrapperSize := 0
	totalRibbonSize := 0
	for _, measurements := range fileContents {
		measures := strings.Split(measurements, "x")
		length, _ := strconv.Atoi(measures[0])
		width, _ := strconv.Atoi(measures[1])
		height, _ := strconv.Atoi(measures[2])
		totalGiftWrapperSize += findGiftSize(length, width, height)
		totalRibbonSize += findRibbonSize(length, width, height)

	}
	fmt.Println(totalGiftWrapperSize)
	fmt.Println(totalRibbonSize)
}

// / find surface are of a right rectangularr prism plus the area with the smallest size
func findGiftSize(length int, width int, height int) int {
	surfaceArea := (2 * length * width) + (2 * width * height) + (2 * length * height)
	giftsize := surfaceArea + findSmallestArea(length, width, height)
	return giftsize
}

func findRibbonSize(length int, width int, height int) int {
	nums := []int{length, width, height}
	sort.Ints(nums)
	return (nums[0] + nums[0]) + (nums[1] + nums[1]) + (length * width * height)
}

// / Given a right rectangularr prism, find the area with the smallest size
func findSmallestArea(length int, width int, height int) int {
	area1 := length * width
	area2 := length * height
	area3 := width * height
	min := area1
	if area2 < min {
		min = area2
	}
	if area3 < min {
		min = area3
	}
	return min
}
