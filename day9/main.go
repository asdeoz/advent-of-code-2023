package main

import (
	h "day9/helpers"
	"fmt"
)

func isLastRow(row []int) bool {
	for _, n := range row {
		if n != 0 {
			return false
		}
	}

	return true
}

func getRowBelow(row []int) []int {
	var below []int

	for i := 0; i < len(row)-1; i++ {
		below = append(below, row[i+1]-row[i])
	}

	return below
}

func getMatrix(row []int) [][]int {
	matrix := [][]int{row}
	current := row

	for {
		if isLastRow(current) {
			break
		}

		below := getRowBelow(current)
		matrix = append(matrix, below)
		current = below
	}

	return matrix
}

func expandMatrix(matrix *[][]int) {
	lastRowInd := len(*matrix) - 1
	(*matrix)[lastRowInd] = append((*matrix)[lastRowInd], 0)

	for i := lastRowInd; i > 0; i-- {
		lastInd := len((*matrix)[i]) - 1
		lastNum := (*matrix)[i][lastInd]
		aboveLastInd := len((*matrix)[i-1]) - 1
		aboveLastNum := (*matrix)[i-1][aboveLastInd]

		(*matrix)[i-1] = append((*matrix)[i-1], lastNum+aboveLastNum)
	}
}

func main() {
	input := h.LoadFile[[][]int]("input.json")
	finalResult := 0
	for _, line := range *input {
		res := getMatrix(line)
		expandMatrix(&res)
		finalResult += res[0][len(res[0])-1]
	}
	fmt.Println(finalResult)
}
