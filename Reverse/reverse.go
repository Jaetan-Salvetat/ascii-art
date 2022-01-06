package Reverse

import (
	"bufio"
	"os"

	"../Justify"
	"../Typo"
)

func Reverse(fileName, typo, justifyValue string, PrintFunc func(string)) {
	var lines []string
	if !ReadLines(fileName, &lines) {
		return
	}

	char := []string{"", "", "", "", "", "", "", ""}
	var toPrint string = ""
	var toAdd string = ""
	incremente := 0

	for lines != nil {
		if incremente >= len(lines[0]) {
			break
		}
		for i := range lines {
			char[i] += string(lines[i][incremente])
		}
		toAdd = Typo.IsStandard(char)
		if toAdd != "" {
			toPrint += toAdd
			char = []string{"", "", "", "", "", "", "", ""}
		}
		incremente++
	}

	if toPrint != "" {
		Justify.Justify(justifyValue, toPrint, typo, true, PrintFunc)
	} else {
		print("not found")
	}
}

func ReadLines(fileName string, lines *[]string) bool {
	file, hasError := os.Open(fileName)

	if hasError != nil {
		print("Erreur de lecture de fichier")
		return false
	}

	var content []string
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		content = append(content, scanner.Text())
	}

	*lines = content
	return true
}
