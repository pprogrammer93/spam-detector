package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const CHAR_NUMBER int = 128
const NOTHING int = -1

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func normalize(text string) string {
	runes := []rune(text)
	for i := 0; i < len(runes); i++ {
		if int(runes[i]) >= CHAR_NUMBER {
			runes[i] = '0'
		}
	}
	return string(runes)
}

func analyzeLastOccurence(keywords string) [CHAR_NUMBER]int {
	var charIndices [CHAR_NUMBER]int

	for i := 0; i < CHAR_NUMBER; i++ {
		charIndices[i] = NOTHING
	}
	for i := 0; i < len(keywords); i++ {
		ascii := keywords[i]
		charIndices[ascii] = i
	}

	return charIndices
}

func solve(keywords, text string) int {
	keywords = strings.ToLower(keywords)
	text = strings.ToLower(normalize(text))

	charIndices := analyzeLastOccurence(keywords)

	jMax := len(keywords) - 1
	iMax := len(text)
	i := jMax
	j := 0
	found := -1

	for i < iMax && found == -1 {
		j = 0
		for j <= jMax && keywords[jMax-j] == text[i-j] {
			j++
		}
		if j > jMax {
			found = i - (j - 1)
		} else {
			indexInKeywords := jMax - j
			indexInText := i - j
			charInText := text[indexInText]
			if charIndices[charInText] == NOTHING {
				i += indexInKeywords + 1
			} else if charIndices[charInText] > indexInKeywords {
				i++
			} else {
				i += indexInKeywords - charIndices[charInText]
			}
		}
	}

	return found
}

func readJSON(filename string) (string, string) {
	type Input struct {
		Keywords string
		Text     string
	}
	var input Input

	inputJSON, ioErr := ioutil.ReadFile(filename)
	check(ioErr)

	jsonErr := json.Unmarshal(inputJSON, &input)
	if jsonErr == nil {
		return input.Keywords, input.Text
	} else {
		return "", ""
	}
}

func main() {
	keywords, text := readJSON(os.Args[1])
	fmt.Println(keywords)
	fmt.Println(text)
	idx := solve(keywords, text)

	if idx != -1 {
		fmt.Println(text[idx:])
	}

	fmt.Println(idx)
}
