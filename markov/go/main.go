package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

var customPunctuation = map[rune]bool{'.': true, '!': true, '?': true}

func buildMarkovModel(text string) map[string][]string {
	words := strings.Fields(text)
	model := make(map[string][]string)

	for i := 0; i < len(words)-1; i++ {
		currentWord := words[i]
		nextWord := words[i+1]

		model[currentWord] = append(model[currentWord], nextWord)
	}

	return model
}

func generateText(model map[string][]string, numSentences int) string {
	rand.Seed(time.Now().UnixNano())

	var generatedText []string
	var currentWord string

	startWordCandidates := make([]string, 0)

	for word := range model {
		if len(word) > 0 && word[0] >= 'A' && word[0] <= 'Z' {
			startWordCandidates = append(startWordCandidates, word)
		}
	}

	if len(startWordCandidates) > 0 {
		currentWord = startWordCandidates[rand.Intn(len(startWordCandidates))]
	} else {
		for word := range model {
			currentWord = word
			break
		}
	}

	sentenceCount := 0

	for sentenceCount < numSentences {
		nextWordOptions, ok := model[currentWord]
		if !ok {
			break
		}

		nextWord := nextWordOptions[rand.Intn(len(nextWordOptions))]
		generatedText = append(generatedText, nextWord)
		currentWord = nextWord

		// Check for the presence of a punctuation mark
		if customPunctuation[rune(nextWord[len(nextWord)-1])] {
			sentenceCount++
		}
	}

	return strings.Join(generatedText, " ")
}

func loadTextFiles(directoryPath string, numFiles int) []string {
	var texts []string

	for i := 1; i <= numFiles; i++ {
		fileName := fmt.Sprintf("%03d.txt", i)
		filePath := fmt.Sprintf("%s/%s", directoryPath, fileName)

		content, err := ioutil.ReadFile(filePath)
		if err == nil {
			texts = append(texts, string(content))
		} else {
			fmt.Printf("File not found: %s\n", filePath)
		}
	}

	return texts
}

func main() {

	directoryPath := "../corpus"
	startTime := time.Now()
	inputTexts := loadTextFiles(directoryPath, 511)
	combinedText := strings.Join(inputTexts, " ")
	markovModel := buildMarkovModel(combinedText)
	generatedText := generateText(markovModel, 500000)
	endTime := time.Now()
	fmt.Println("Generated Text:")
	fmt.Println(generatedText)
	executionTime := endTime.Sub(startTime)
	fmt.Printf("Execution Time: %f seconds\n", executionTime.Seconds())
	fmt.Println("Press Enter to exit...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
	
}
