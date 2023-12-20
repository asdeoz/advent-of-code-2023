package h

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func LoadFile[T any](fileName string) *T {
	var loadedFile T
	file, err := os.Open(fileName)
	defer file.Close()
	if err != nil {
		fmt.Println("Error reading", fileName, ":", err)
	}

	inputByteValue, _ := io.ReadAll(file)

	json.Unmarshal(inputByteValue, &loadedFile)

	return &loadedFile
}

func Ternary[T any](comp bool, value1 T, value2 T) T {
	if comp {
		return value1
	} else {
		return value2
	}
}

func gcd(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

// find Least Common Multiple (LCM) via GCD
func FindLCM(a, b int, integers ...int) int {
	result := a * b / gcd(a, b)

	for i := 0; i < len(integers); i++ {
		result = FindLCM(result, integers[i])
	}

	return result
}
