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
		winningNumbers := c.SumWinningNumbers()
		if winningNumbers > 0 {
			sum += int(math.Pow(2, float64(winningNumbers-1)))
		}
	}
	return sum
}

func getNumberOfCards(cards *[]models.ScratchCard, ogCards *[]models.ScratchCard, startingIndex int) int {
	totalNumOfCards := len(*cards)

	for index, c := range *cards {
		winningNums := c.SumWinningNumbers()
		if winningNums > 0 {
			nextCards := (*ogCards)[startingIndex+index+1 : startingIndex+index+1+winningNums]
			totalNumOfCards += getNumberOfCards(&nextCards, ogCards, startingIndex+index+1)
		}
	}

	return totalNumOfCards
}

func main() {
	cards := readCards()

	part1Result := calculatePointsFromCards(cards)

	fmt.Println("Result of Part 1: ", part1Result)

	part2Result := getNumberOfCards(cards, cards, 0)

	fmt.Println("Result of Part 2: ", part2Result)
}
