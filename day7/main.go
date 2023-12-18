package main

import (
	"day7/models"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"regexp"
	"sort"
	"sync"
)

func getInput() *[]models.Deal {
	var deals []models.Deal
	inputFile, err := os.Open("input.json")
	// inputFile, err := os.Open("example.json")
	defer inputFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	inputByteValue, _ := io.ReadAll(inputFile)

	json.Unmarshal(inputByteValue, &deals)

	return &deals
}

func lenRe(re *regexp.Regexp, hand string) int {
	return len(re.FindAllString(hand, 5))
}

func findType(hand string) int {
	findings := make([]int, 0, 13)

	re2 := regexp.MustCompile(`2`)
	re3 := regexp.MustCompile(`3`)
	re4 := regexp.MustCompile(`4`)
	re5 := regexp.MustCompile(`5`)
	re6 := regexp.MustCompile(`6`)
	re7 := regexp.MustCompile(`7`)
	re8 := regexp.MustCompile(`8`)
	re9 := regexp.MustCompile(`9`)
	reT := regexp.MustCompile(`T`)
	reJ := regexp.MustCompile(`J`)
	reQ := regexp.MustCompile(`Q`)
	reK := regexp.MustCompile(`K`)
	reA := regexp.MustCompile(`A`)

	findings = append(
		findings,
		lenRe(re2, hand),
		lenRe(re3, hand),
		lenRe(re4, hand),
		lenRe(re5, hand),
		lenRe(re6, hand),
		lenRe(re7, hand),
		lenRe(re8, hand),
		lenRe(re9, hand),
		lenRe(reT, hand),
		lenRe(reJ, hand),
		lenRe(reQ, hand),
		lenRe(reK, hand),
		lenRe(reA, hand),
	)

	found5, found4, foundFull, foundThree, foundTwoPair, foundOnePair := false, false, false, false, false, false

	for _, f := range findings {
		if f == 5 {
			found5 = true
		}
		if f == 4 {
			found4 = true
		}
		if f == 3 {
			if foundOnePair {
				foundFull = true
			}
			foundThree = true
		}
		if f == 2 {
			if foundOnePair {
				foundTwoPair = true
			}
			if foundThree {
				foundFull = true
			}
			foundOnePair = true
		}
	}

	if found5 {
		return 1
	}
	if found4 {
		return 2
	}
	if foundFull {
		return 3
	}
	if foundThree {
		return 4
	}
	if foundTwoPair {
		return 5
	}
	if foundOnePair {
		return 6
	}
	return 7
}

func assignTypes(deals *[]models.Deal) {
	var wg sync.WaitGroup

	for i := 0; i < len(*deals); i++ {
		wg.Add(1)
		go func(deal *models.Deal) {
			defer wg.Done()
			(*deal).Type = findType((*deal).Hand)
		}(&((*deals)[i]))
	}

	wg.Wait()
}

func getLetterValue(letter string) int {
	switch letter {
	case "2":
		return 2
	case "3":
		return 3
	case "4":
		return 4
	case "5":
		return 5
	case "6":
		return 6
	case "7":
		return 7
	case "8":
		return 8
	case "9":
		return 9
	case "T":
		return 10
	case "J":
		return 11
	case "Q":
		return 12
	case "K":
		return 13
	case "A":
		return 14
	}

	return 0
}

func isHand1BetterThanHand2(hand1, hand2 string) bool {
	if hand1[0:1] != hand2[0:1] {
		return getLetterValue(hand1[0:1]) > getLetterValue(hand2[0:1])
	}
	if hand1[1:2] != hand2[1:2] {
		return getLetterValue(hand1[1:2]) > getLetterValue(hand2[1:2])
	}
	if hand1[2:3] != hand2[2:3] {
		return getLetterValue(hand1[2:3]) > getLetterValue(hand2[2:3])
	}
	if hand1[3:4] != hand2[3:4] {
		return getLetterValue(hand1[3:4]) > getLetterValue(hand2[3:4])
	}
	if hand1[4:5] != hand2[4:5] {
		return getLetterValue(hand1[4:5]) > getLetterValue(hand2[4:5])
	}

	return false
}

func sortDeals(deals *[]models.Deal) {
	sort.Slice(*deals, func(i, j int) bool {
		if (*deals)[i].Type == (*deals)[j].Type {
			return isHand1BetterThanHand2((*deals)[i].Hand, (*deals)[j].Hand)
		}
		return (*deals)[i].Type < (*deals)[j].Type
	})
}

func assignRanks(deals *[]models.Deal) {
	lenDeals := len(*deals)
	for i := 0; i < lenDeals; i++ {
		(*deals)[i].Rank = lenDeals - i
	}
}

func calcWinnings(deals *[]models.Deal) int {
	totalWinnings := 0
	for _, d := range *deals {
		totalWinnings += d.Bid * d.Rank
	}
	return totalWinnings
}

func main() {
	deals := getInput()

	assignTypes(deals)

	sortDeals(deals)

	assignRanks(deals)

	winnings := calcWinnings(deals)

	fmt.Println("Winnings of Part 1:", winnings)
}
