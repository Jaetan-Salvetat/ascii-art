package Typo

import (
	"bufio"
	"io/ioutil"
	"os"
)

func AskChar(typo string, char rune) []string {
	var toReturn []string
	char -= 32

	files := ListFile()
	var lines []string

	for _, val := range files {
		if val == typo {
			val = "./Typo/" + val + ".txt"
			if !ReadLines(val, &lines) {
				return toReturn
			}
		}
	}
	if lines == nil {
		print("aucun fichier correspondant à la typo demandée")
	}

	toReturn = GetTxtChar(int(char), lines)

	return toReturn
}

func ListFile() []string {
	files := []string{""}

	fileList, _ := ioutil.ReadDir("./Typo")

	for i := range fileList {
		val := fileList[i].Name()
		if len(val) > 4 {
			if val[len(val)-4:] == ".txt" {
				files = append(files, val[:len(val)-4])
			}
		}
	}
	return files
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

func GetTxtChar(char int, lines []string) []string {
	var toReturn []string
	incremente := 0
	for i := 1; i < len(lines); i++ {
		if lines[i] == "" {
			incremente++
		}
		if incremente == char {
			toReturn = lines[i+1 : i+10]
			break
		}
	}
	return toReturn
}
