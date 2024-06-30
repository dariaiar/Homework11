package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {
	var fileName string
	var expression string

	fileName = "listofnumbers.txt"
	expression = `^\d{10}|\(\d{3}\) \d{3}-\d{4}|\(\d{3}\)\d{3}-\d{4}|\d{3}\-\d{3}-\d{4}|\d{3}.\d{3}.\d{4}|\d{3}\ \d{3}\ \d{4}|\(\d{3}\)\d{7}|\(\d{3}\) \d{7}`
	findTelNumber(fileName, expression)

	fileName = "text.txt"
	expression = `(\b.*ти\b)|(\b.*ть\b)|(\b.*ся\b)|(\b.*сь\b)`
	//Знайти всі слова неозначеної форми дієслова(інфінітив). Закінчення -ти,-ть,-ся,-сь
	findTelNumber(fileName, expression)
}

func findTelNumber(fileName string, expression string) {
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Printf("reading file: %v", err)
		return
	}

	pattern := regexp.MustCompile(expression)
	lines := strings.Split(string(fileContent), "\n")
	for b, line := range lines {
		fmt.Printf("Line %v: %v\n", b, line)
		results := pattern.FindAllString(line, -1)
		for i, v := range results {
			fmt.Printf("Match %v: %v\n", i, v)
		}
	}
}
