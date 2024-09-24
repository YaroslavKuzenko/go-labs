package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

func shuffleWord(word string) string {
	rand.Seed(time.Now().UnixNano())
	runes := []rune(word)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}

func replaceDoubleLetters(word string) string {
	var result []rune
	for i := 0; i < len(word); i++ {
		if i > 0 && word[i] == word[i-1] {
			result = append(result, '+')
		} else {
			result = append(result, rune(word[i]))
		}
	}
	return string(result)
}

func main() {
	inputFile := "input.txt"

	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Помилка при читанні файлу:", err)
		return
	}

	lines := strings.Split(string(content), "\n")
	var processedLines []string

	for _, line := range lines {
		words := strings.Fields(line)
		for i, word := range words {
			words[i] = replaceDoubleLetters(word)
			words[i] = shuffleWord(words[i])
		}
		processedLines = append(processedLines, strings.Join(words, "-"))
	}

	result := strings.Join(processedLines, "\n")

	outputFile := "output.txt"
	err = ioutil.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Помилка при записі у файл:", err)
		return
	}

	fmt.Println("Обробка завершена. Результат записано у", outputFile)
}
