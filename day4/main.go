package main

import (
	"day4/models"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
)

func readCards() *[]models.ScratchCard {
	var input []models.ScratchCard

	jsonFile, err := os.Open("input.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &input)

	return &input
}

func calculatePointsFromCards(cards *[]models.ScratchCard) int {
	sum := 0
	for _, c := range *cards {
		winningNumbers := 0
		for _, n := range c.Nums {
			if c.IsWinningNumber(n) {
				winningNumbers++
			}
		}
		if winningNumbers > 0 {
			sum += int(math.Pow(2, float64(winningNumbers-1)))
		}
	}
	return sum
}

func main() {
	cards := readCards()

	part1Result := calculatePointsFromCards(cards)

	fmt.Println("Result of Part 1: ", part1Result)
}
