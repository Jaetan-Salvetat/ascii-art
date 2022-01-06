package Justify

func NormalPrint(s []string, Print func(string)) {
	for _, val := range s { //Print the sentense characters one by one
		Print(val)
		Print("\n")
	}
}

func FuseTable(table, toAdd []string) []string { //add letter table to full sentense table
	for i := range table {
		table[i] += toAdd[i]
	}
	return table
}
