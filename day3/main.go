package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

const NUM_LINES = 140
const LENGTH_LINE = 140

func getMatrix() [][]string {
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

func ternary[T any](condition bool, res1 T, res2 T) T {
	if condition {
		return res1
	}
	return res2
}

func checkBoundary(positions *[]int, line int, matrix *[][]string) bool {
	pos := *positions
	r := regexp.MustCompile(`[^\.\d]`)
	lastPos := len(pos) - 1
	lineAbove, lineBelow, charBefore, charAfter := "", "", "", ""
	lineFirstPos := ternary(pos[0] == 0, 0, pos[0]-1)
	lineLastPos := ternary(pos[lastPos] == LENGTH_LINE-1, pos[lastPos], pos[lastPos]+1)

	if line != 0 {
		lineAbove = strings.Join((*matrix)[line-1][lineFirstPos:lineLastPos+1], "")
	}
	if line != NUM_LINES-1 {
		lineBelow = strings.Join((*matrix)[line+1][lineFirstPos:lineLastPos+1], "")
	}
	if pos[0] > 0 {
		charBefore = (*matrix)[line][pos[0]-1]
	}
	if pos[lastPos] < LENGTH_LINE-1 {
		charAfter = (*matrix)[line][pos[lastPos]+1]
	}

	return r.MatchString(lineAbove + lineBelow + charBefore + charAfter)
}

func toIntArray(array *[]string) []int {
	result := make([]int, len(*array))
	for i := 0; i < len(*array); i++ {
		result[i], _ = strconv.Atoi((*array)[i])
	}
	return result
}

func toNumber(array *[]string) int {
	numbers := toIntArray(array)
	slices.Reverse(numbers)
	result := 0
	for i := 0; i < len(numbers); i++ {
		result += numbers[i] * int(math.Pow(10.0, float64(i)))
	}
	return result
}

func traverseMatrix(matrix *[][]string) []int {
	r := regexp.MustCompile(`\d`)
	result := make([]int, 10000)
	buffer := make([]string, 0, 10)
	positions := make([]int, 0, 10)

	for i := 0; i < len(*matrix); i++ {
		currLine := (*matrix)[i]
		for j := 0; j < len(currLine); j++ {
			if r.MatchString(currLine[j]) {
				buffer = append(buffer, currLine[j])
				positions = append(positions, j)
			}
			if (!r.MatchString(currLine[j]) || j == len(currLine)-1) && len(positions) > 0 {
				if checkBoundary(&positions, i, matrix) {
					validNumber := toNumber(&buffer)
					result = append(result, validNumber)
				}
				buffer = make([]string, 0, 10)
				positions = make([]int, 0, 10)
			}
		}
	}

	return result
}

func main() {
	matrix := getMatrix()

	result := traverseMatrix(&matrix)
	sum := 0
	for _, value := range result {
		sum += value
	}
	fmt.Println(sum)
}
