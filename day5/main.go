package main

import (
	"day5/models"
	"encoding/json"
	"fmt"
	"io"
	"math"
	"os"
	"sync"
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

func expandSeedsAndFind(seeds *[]int, input *models.Map) []int {
	var wg sync.WaitGroup
	locations := make([]int, 0, 100000000)

	n1 := 0
	for _, n := range *seeds {
		if n1 == 0 {
			n1 = n
		} else {
			wg.Add(1)
			go func(n1 int, n2 int, input *models.Map, locations *[]int) {
				defer wg.Done()
				lowestLocation := math.MaxInt32
				for i := n1; i < n2; i++ {
					loc := findLocation(input, i)
					if loc < lowestLocation {
						lowestLocation = loc
					}
				}
				*locations = append(*locations, lowestLocation)
			}(n1, n1+n, input, &locations)
			n1 = 0
		}
	}

	wg.Wait()
	return locations
}

func main() {
	// Initialize
	input, seeds := readInput()

	// Part 1
	lowestLocation1 := math.MaxInt32
	for _, s := range *seeds {
		loc := findLocation(input, s)
		if loc < lowestLocation1 {
			lowestLocation1 = loc
		}
	}
	fmt.Println("Lowest Location for Part 1:", lowestLocation1)

	// Part 2
	locations := expandSeedsAndFind(seeds, input)
	lowestLocation2 := math.MaxInt32
	for _, l := range locations {
		if l < lowestLocation2 {
			lowestLocation2 = l
		}
	}

	fmt.Println("Lowest Location for Part 2:", lowestLocation2)
}
