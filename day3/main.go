package main

import (
	"day3/models"
	"day3/utils"
	"fmt"
	"regexp"
	"strings"
)

const NUM_LINES = 140
const LENGTH_LINE = 140

// Part 1
func checkBoundary(positions *[]int, line int, matrix *[][]string) bool {
	pos := *positions
	r := regexp.MustCompile(`[^\.\d]`)
	lastPos := len(pos) - 1
	lineAbove, lineBelow, charBefore, charAfter := "", "", "", ""
	lineFirstPos := utils.Ternary(pos[0] == 0, 0, pos[0]-1)
	lineLastPos := utils.Ternary(pos[lastPos] == LENGTH_LINE-1, pos[lastPos], pos[lastPos]+1)

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

func traverseMatrixPart1(matrix *[][]string) []int {
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
					validNumber := utils.ToNumber(&buffer)
					result = append(result, validNumber)
				}
				buffer = make([]string, 0, 10)
				positions = make([]int, 0, 10)
			}
		}
	}

	return result
}

func calculatePart1(matrix *[][]string) int {
	result := traverseMatrixPart1(matrix)
	sum := 0
	for _, value := range result {
		sum += value
	}
	return sum
}

// Part 2
func getGearIds(matrix *[][]string, lineNum int, first int, last int) []string {
	ids := make([]string, 0, 10)

	for i := first; i <= last; i++ {
		if (*matrix)[lineNum][i] == "*" {
			ids = append(ids, fmt.Sprint(lineNum, "_", i))
		}
	}

	return ids
}

func getSurroundingGears(positions *[]int, line int, matrix *[][]string, number int) *[]models.Gear {
	pos := *positions
	lastPos := len(pos) - 1
	lineFirstPos := utils.Ternary(pos[0] == 0, 0, pos[0]-1)
	lineLastPos := utils.Ternary(pos[lastPos] == LENGTH_LINE-1, pos[lastPos], pos[lastPos]+1)

	var lineAbove, lineBelow, charBefore, charAfter []string

	if line != 0 {
		lineAbove = getGearIds(matrix, line-1, lineFirstPos, lineLastPos)
	}
	if line != NUM_LINES-1 {
		lineBelow = getGearIds(matrix, line+1, lineFirstPos, lineLastPos)
	}
	if pos[0] > 0 {
		charBefore = getGearIds(matrix, line, pos[0]-1, pos[0]-1)
	}
	if pos[lastPos] < LENGTH_LINE-1 {
		charAfter = getGearIds(matrix, line, pos[lastPos]+1, pos[lastPos]+1)
	}

	resultSize := len(lineAbove) + len(lineBelow) + len(charBefore) + len(charAfter)
	allIds := append(lineAbove, lineBelow...)
	allIds = append(allIds, charBefore...)
	allIds = append(allIds, charAfter...)

	result := make([]models.Gear, resultSize)
	for i := 0; i < len(allIds); i++ {
		result[i] = models.Gear{Id: allIds[i], Number1: number}
	}

	return &result
}

func traverseMatrixPart2(matrix *[][]string) int {
	r := regexp.MustCompile(`\d`)
	// result := make([]int, 10000)
	buffer := make([]string, 0, 10)
	positions := make([]int, 0, 10)
	gears := make([]models.Gear, 0, 10000)

	for i := 0; i < len(*matrix); i++ {
		currLine := (*matrix)[i]
		for j := 0; j < len(currLine); j++ {
			if r.MatchString(currLine[j]) {
				buffer = append(buffer, currLine[j])
				positions = append(positions, j)
			}
			if (!r.MatchString(currLine[j]) || j == len(currLine)-1) && len(positions) > 0 {
				number := utils.ToNumber(&buffer)
				surroundingGears := getSurroundingGears(&positions, i, matrix, number)
				gears = append(gears, *surroundingGears...)
				buffer = make([]string, 0, 10)
				positions = make([]int, 0, 10)
			}
		}
	}

	for i := 0; i < len(gears); i++ {
		for j := i + 1; j < len(gears); j++ {
			if gears[i].Id == gears[j].Id {
				gears[i].Number2 = gears[j].Number1
			}
		}
	}

	sum := 0
	for _, value := range gears {
		sum += value.Number1 * value.Number2
	}
	return sum
}

func main() {
	matrix := utils.GetMatrix()

	part1Result := calculatePart1(&matrix)
	fmt.Println("Result of Part 1: ", part1Result)

	part2Result := traverseMatrixPart2(&matrix)
	fmt.Println("Result of Part 2: ", part2Result)
}
