package Justify

import (
	"github.com/nathan-fiscaletti/consolesize-go"
)

func Justify(align, text, typo string, isReverse bool, Print func(string)) {
	cols, _ := consolesize.GetConsoleSize()
	toPrint := []string{
		"",
		"",
		"",
		"",
		"",
		"",
		"",
		"",
	}

	switch align {
	case "left":
		left(text, typo, cols, isReverse, toPrint, Print)
	case "right":
		right(text, typo, cols, isReverse, toPrint, Print)
	case "center":
		center(text, typo, cols, isReverse, toPrint, Print)
	case "justify":
		justify(text, typo, cols, isReverse, toPrint, Print)
	default:
	}
}
