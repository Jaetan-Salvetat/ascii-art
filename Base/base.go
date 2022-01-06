package Base

import (
	"../Colors"
	"../Justify"
	"../Output"
	"../Reverse"
)

type Params struct {
	useColor  bool
	output    bool
	reverse   bool
	justify   bool
	fontStyle bool
	hasText   bool

	colorValue      string
	outputFileName  string
	justifyValue    string
	fontStyleValue  string
	reverseFileName string
	text            string
}

func PrintASCII(args []string) {

	params := Params{justifyValue: "left", fontStyleValue: "standard"}

	if !VerifArgs(args, &params) {
		return
	}

	PrintFunc := Colors.GetColor(params.colorValue)

	if params.output {
		Output.Output(params.outputFileName, params.text, params.fontStyleValue)
		return
	}
	if params.reverse {
		Reverse.Reverse(params.reverseFileName, params.fontStyleValue, params.justifyValue, PrintFunc)
		return
	}
	if params.justify {
		Justify.Justify(params.justifyValue, params.text, params.fontStyleValue, false, PrintFunc)
		return
	}
	NormalPrint(params, PrintFunc)
}
