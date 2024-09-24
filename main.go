package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strings"
	"time"
)

// Функція для перемішування букв у слові
func shuffleWord(word string) string {
	rand.Seed(time.Now().UnixNano())
	runes := []rune(word)
	rand.Shuffle(len(runes), func(i, j int) {
		runes[i], runes[j] = runes[j], runes[i]
	})
	return string(runes)
}

// Функція для перевірки подвоєних літер і заміни їх на "+"
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
	// Вхідний файл
	inputFile := "input.txt"

	// Зчитуємо вміст файлу
	content, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println("Помилка при читанні файлу:", err)
		return
	}

	// Розділяємо на рядки
	lines := strings.Split(string(content), "\n")
	var processedLines []string

	for _, line := range lines {
		// Розділяємо на слова
		words := strings.Fields(line)
		for i, word := range words {
			// Заміна подвоєних літер
			words[i] = replaceDoubleLetters(word)
			// Перемішування букв у слові
			words[i] = shuffleWord(words[i])
		}
		// Об'єднуємо слова назад у рядок
		processedLines = append(processedLines, strings.Join(words, " "))
	}

	// Результуючий рядок
	result := strings.Join(processedLines, "\n")

	// Записуємо результат у новий файл
	outputFile := "output.txt"
	err = ioutil.WriteFile(outputFile, []byte(result), 0644)
	if err != nil {
		fmt.Println("Помилка при записі у файл:", err)
		return
	}

	fmt.Println("Обробка завершена. Результат записано у", outputFile)
}
