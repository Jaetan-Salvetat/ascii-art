package Output

import (
	"os"

	"../Typo"
)

func Output(fileName, s, typo string) {
	var result []string
	var temp, res string
	fs, err := os.Create(fileName)

	if err != nil {
		print(err)
		return
	}

	for i := 0; i < len(s); i++ {
		x := s[i]
		if rune(x) == '\\' && s[i+1] == 'n' {
			result = append(result, temp)
			i++
			temp = ""
		} else {
			temp += string(x)
		}
	}
	if temp != "" {
		result = append(result, temp)
	}

	for _, v := range result {
		for i := 0; i < 8; i++ {
			for _, val := range v {
				toAdd := Typo.AskChar(typo, val)

				for _, val := range toAdd[i] {
					res += string(val)
				}
			}
			res += "\n"
		}
	}

	fs.WriteString(res)
}
