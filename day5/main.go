package main

import (
	"day5/models"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
)

func readInput() (*models.Map, *[]int) {
	var input models.Map
	var seeds []int

	inputFile, err := os.Open("input.json")
	// inputFile, err := os.Open("example.json")
	defer inputFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	inputByteValue, _ := io.ReadAll(inputFile)

	json.Unmarshal(inputByteValue, &input)

	seedsFile, err := os.Open("seeds.json")
	// seedsFile, err := os.Open("example-seeds.json")
	defer seedsFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	seedsByteValue, _ := io.ReadAll(seedsFile)

	json.Unmarshal(seedsByteValue, &seeds)

	return &input, &seeds
}

func findNext(ranges []models.Range, nextIndex int) int {
	for _, r := range ranges {
		if nextIndex >= r.SourceStart && nextIndex < r.SourceStart+r.RangeLen {
			return (nextIndex - r.SourceStart) + r.DestStart
		}
	}

	return nextIndex
}

func findLocation(m *models.Map, seed int) int {
	soil := findNext(m.SeedToSoil, seed)
	fertilizer := findNext(m.SoilToFertilizer, soil)
	water := findNext(m.FertilizerToWater, fertilizer)
	light := findNext(m.WaterToLight, water)
	temperature := findNext(m.LightToTemperature, light)
	humidity := findNext(m.TemperatureToHumidity, temperature)
	location := findNext(m.HumidityToLocation, humidity)

	return location
}

func main() {
	input, seeds := readInput()
	lowestLocation := math.MaxInt32
	for _, s := range *seeds {
		loc := findLocation(input, s)
		if loc < lowestLocation {
			lowestLocation = loc
		}
	}
	fmt.Println("Lowest Location: ", lowestLocation)
}
