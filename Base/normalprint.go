package Base

import "../Typo"

func NormalPrint(params Params, PrintFunc func(string)) {
	var toAdd []string
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

	for i := 0; i < len(params.text); i++ {
		if i != len(params.text)-1 {
			if params.text[i] == '\\' && params.text[i+1] == 'n' {
				if toPrint[1] != "" {
					PrintLine(toPrint, PrintFunc)
					toPrint = []string{"", "", "", "", "", "", "", ""}
					i += 2
				} else {
					i += 2
				}
			}
		}
		//toAdd = Typo.GetChar(params.fontStyleValue, string(val)) // Fonction qui utilise les fichiers GO pour la police
		toAdd = Typo.AskChar(params.fontStyleValue, rune(params.text[i])) // Fonction qui utilise les fichier textes.
		if toAdd != nil {
			toPrint = FuseTable(toPrint, toAdd)
		}
	}
	if toPrint[1] != "" {
		PrintLine(toPrint, PrintFunc)
	}
}

func FuseTable(table, toAdd []string) []string { //add letter table to full sentense table
	for i := range table {
		table[i] += toAdd[i]
	}
	return table
}

func PrintLine(toPrint []string, PrintFunc func(string)) {
	for _, val := range toPrint { //Print the sentense characters one by one
		PrintFunc(val)
		PrintFunc("\n")
	}
}
