package utils

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
)

func GetMatrix() [][]string {
	txtFile, err := os.Open("input.txt")
	defer txtFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	matrix := make([][]string, 140)
	curr := 0

	scanner := bufio.NewScanner(txtFile)
	for scanner.Scan() {
		line := scanner.Text()
		splitLine := make([]string, len(line))
		for i := 0; i < len(line); i++ {
			splitLine[i] = string(line[i])
		}
		matrix[curr] = splitLine
		curr++
	}

	return matrix
}

func Ternary[T any](condition bool, res1 T, res2 T) T {
	if condition {
		return res1
	}
	return res2
}

func ToIntArray(array *[]string) []int {
	result := make([]int, len(*array))
	for i := 0; i < len(*array); i++ {
		result[i], _ = strconv.Atoi((*array)[i])
	}
	return result
}

func ToNumber(array *[]string) int {
	numbers := ToIntArray(array)
	slices.Reverse(numbers)
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i] * int(math.Pow(10.0, float64(i)))
	}
	return result
}
