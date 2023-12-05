package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type Hand struct {
	R *int `json:"r,omitempty"`
	G *int `json:"g,omitempty"`
	B *int `json:"b,omitempty"`
}

type Game struct {
	Id    int    `json:"id"`
	Hands []Hand `json:"hands"`
}

/*
Determine which games would have been possible if the bag had been
loaded with only 12 red cubes, 13 green cubes, and 14 blue cubes.
What is the sum of the IDs of those games?
*/
func findPossibleGamesSum(input *[]Game) int {
	red := 12
	green := 13
	blue := 14

	result := 0

	for i := 0; i < len(*input); i++ {
		greenNil, redNil, blueNil := true, true, true
		valid := true

		for j := 0; j < len((*input)[i].Hands); j++ {
			hand := (*input)[i].Hands[j]
			if hand.B != nil {
				blueNil = false
				if *hand.B > blue {
					valid = false
				}
			}
			if hand.R != nil {
				redNil = false
				if *hand.R > red {
					valid = false
				}
			}
			if hand.G != nil {
				greenNil = false
				if *hand.G > green {
					valid = false
				}
			}
		}

		if valid && !blueNil && !greenNil && !redNil {
			result += (*input)[i].Id
		}
	}

	return result
}

/*
For each game, find the minimum set of cubes that must have been present.
What is the sum of the power of these sets?
*/
func findMinimumGamesPower(input *[]Game) int {
	result := 0

	for i := 0; i < len(*input); i++ {
		hands := (*input)[i].Hands
		minG, minB, minR := 0, 0, 0

		for j := 0; j < len(hands); j++ {
			hand := hands[j]
			if hand.G != nil && *hand.G > minG {
				minG = *hand.G
			}
			if hand.R != nil && *hand.R > minR {
				minR = *hand.R
			}
			if hand.B != nil && *hand.B > minB {
				minB = *hand.B
			}
		}

		result += (minB * minG * minR)
	}

	return result
}

func main() {
	var input []Game

	jsonFile, err := os.Open("input.json")
	defer jsonFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened input.json")

	byteValue, _ := io.ReadAll(jsonFile)

	json.Unmarshal(byteValue, &input)

	possibleGamesResult := findPossibleGamesSum(&input)

	fmt.Println("Possible Games Result: ", possibleGamesResult)

	minGamesPower := findMinimumGamesPower(&input)

	fmt.Println("Minimum Games Power: ", minGamesPower)
}
