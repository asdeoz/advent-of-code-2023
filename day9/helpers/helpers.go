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
