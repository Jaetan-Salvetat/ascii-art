package Justify

import "../Typo"

func left(s, typo string, cols int, isReverse bool, toPrint []string, Print func(string)) {
	if isReverse {
		Print(s)
		return
	}
	var toAdd []string

	for i := 0; i < len(s); i++ {
		if i < len(s)-1 && s[i] == '\\' {
			if s[i+1] == 'n' {
				if i < len(s)-2 {
					NormalPrint(toPrint, Print)
					toPrint = []string{"", "", "", "", "", "", "", ""}
				}
				i++
				continue
			}
		}
		toAdd = Typo.AskChar(typo, rune(s[i]))
		toPrint = FuseTable(toPrint, toAdd)
	}
	NormalPrint(toPrint, Print)
}

func right(s, typo string, cols int, isReverse bool, toPrint []string, Print func(string)) {
	if isReverse {
		var nbrSpace int = cols - len(s) - 2
		var temp string
		for i := 0; i <= nbrSpace; i++ {
			temp += " "
		}
		Print(temp + s)
		return
	}
	var toAdd []string

	for _, val := range s {
		toAdd = Typo.AskChar(typo, val)
		toPrint = FuseTable(toPrint, toAdd)
	}

	if len(toPrint[0]) < cols {
		var temp string
		for j := 0; j < (cols-len(toPrint[0]))-2; j++ {
			temp += " "
		}
		for i := 0; i < len(toPrint); i++ {
			toPrint[i] = temp + toPrint[i]
		}
	}
	NormalPrint(toPrint, Print)
}

func center(s, typo string, cols int, isReverse bool, toPrint []string, Print func(string)) {
	if isReverse {
		var nbrSpace int = (cols - len(s) - 2) / 2
		var temp string
		for i := 0; i <= nbrSpace; i++ {
			temp += " "
		}
		Print(temp + s + temp)
		return
	}
	var toAdd []string

	for _, val := range s {
		toAdd = Typo.AskChar(typo, val)
		toPrint = FuseTable(toPrint, toAdd)
	}

	if len(toPrint[0]) < cols {
		var temp string
		for j := 0; j < (cols-len(toPrint[0])-2)/2; j++ {
			temp += " "
		}
		for i := 0; i < len(toPrint); i++ {
			toPrint[i] = temp + toPrint[i] + temp
		}
	}
	NormalPrint(toPrint, Print)
}

func justify(s, typo string, cols int, isReverse bool, toPrint []string, Print func(string)) {
	var nbrSpace int = 0

	for i, val := range s {
		if string(val) == " " && i < len(s)-1 {
			nbrSpace++
		}
	}

	if isReverse {
		var tempo string
		var space int = (cols - len(s) - 2) / nbrSpace
		for i := 0; i <= space; i++ {
			tempo += " "
		}
		for _, val := range s {
			if string(val) == " " {
				Print(tempo)
			} else {
				Print(string(val))
			}
		}
		return
	}
	var toAdd []string
	temp := []string{"", "", "", "", "", "", "", ""}
	result := []string{"", "", "", "", "", "", "", ""}

	for _, val := range s {
		toAdd = Typo.AskChar(typo, val)
		toPrint = FuseTable(toPrint, toAdd)
	}

	if nbrSpace == 0 {
		nbrSpace = 1
	}

	for i := range temp {
		for j := 0; j < (cols-len(toPrint[0]))/nbrSpace; j++ {
			temp[i] += " "
		}
	}

	for i, val := range s {
		if string(val) == " " && i < len(s)-1 {
			result = FuseTable(result, temp)
		} else {
			toAdd = Typo.AskChar(typo, val)
			toPrint = FuseTable(result, toAdd)
		}
	}

	NormalPrint(result, Print)
}
