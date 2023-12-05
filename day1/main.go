package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type Input struct {
	Input []string `json:"input"`
}

type KeyValue struct {
	Key   *regexp.Regexp
	Value string
}

func getNumbersInLine(line string, r *regexp.Regexp) int {
	number := r.ReplaceAllString(line, "")
	numLen := len(number)
	var curr string
	if numLen > 1 {
		curr = number[0:1] + number[numLen-1:numLen]
	} else {
		curr = number + number
	}
	newNum, err := strconv.Atoi(curr)
	if err != nil {
		panic("Fuck!")
	}

	return newNum
}

func numberCompiler(input *Input) int {
	r := regexp.MustCompile(`\D`)

	result := 0

	for i := 0; i < len(input.Input); i++ {
		result += getNumbersInLine(input.Input[i], r)
	}

	return result
}

func stringCompiler(input *Input) int {
	r := regexp.MustCompile(`\D`)

	code := []KeyValue{
		{Key: regexp.MustCompile("one"), Value: "o1ne"},
		{Key: regexp.MustCompile("two"), Value: "t2wo"},
		{Key: regexp.MustCompile("three"), Value: "t3hree"},
		{Key: regexp.MustCompile("four"), Value: "f4our"},
		{Key: regexp.MustCompile("five"), Value: "f5ive"},
		{Key: regexp.MustCompile("six"), Value: "s6ix"},
		{Key: regexp.MustCompile("seven"), Value: "s7even"},
		{Key: regexp.MustCompile("eight"), Value: "e8ight"},
		{Key: regexp.MustCompile("nine"), Value: "n9ine"},
	}

	result := 0

	for i := 0; i < len(input.Input); i++ {
		line := strings.ToLower(input.Input[i])
		for j := 0; j < len(code); j++ {
			line = code[j].Key.ReplaceAllString(line, code[j].Value)
		}
		number := getNumbersInLine(line, r)
		result += number
	}

	return result
}

func main() {
	// Open our jsonFile
	jsonFile, err := os.Open("input.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()
	// if we os.Open returns an error then handle it
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened input.json")

	byteValue, _ := io.ReadAll(jsonFile)

	var input Input

	json.Unmarshal(byteValue, &input)

	numberResult := numberCompiler(&input)

	fmt.Println("Result for numbers: ", numberResult)

	stringResult := stringCompiler(&input)

	fmt.Println("Result for strings: ", stringResult)
}
