package Base

func VerifArgs(args []string, params *Params) bool {

	if args[0] == "" { //Contiens du texte a imprimer
		return false
	}

	var parts []string

	for _, val := range args { //Vérifie si les arguments sont valides
		parts = []string{}

		if !Parse(val, &parts) {

			if parts[0] == "standard" || parts[0] == "shadow" || parts[0] == "thinkertoy" {
				if params.fontStyle {
					return Error("Argument en doublon")
				}
				params.fontStyleValue = parts[0]
				params.fontStyle = true
			} else {
				if params.hasText {
					println("Argument Inconnu")
				} else {
					params.hasText = true
					params.text = parts[0]
				}
			}
		}

		switch parts[0] { // Dupe error handler and bools init
		case "color":
			if TestColorValues(parts[1]) {
				if params.useColor {
					return Error("Couleur définie deux fois")
				}
				params.colorValue = parts[1]
				params.useColor = true
			} else {
				return Error("Couleur non valide: " + parts[1])
			}
		case "output":
			if parts[1][len(parts[1])-4:] == ".txt" && len(parts[1]) > 4 {
				if params.output {
					return Error("Output défini deux fois")
				}
				params.output = true
				params.outputFileName = parts[1]
			} else {
				return Error("Nom de fichier incorrect: " + parts[1])
			}
		case "align":
			if parts[1] == "center" || parts[1] == "left" || parts[1] == "right" || parts[1] == "justify" {
				if params.justify {
					return Error("Argument en doublon")
				}
				params.justify = true
				params.justifyValue = parts[1]
			}
		case "reverse":
			if parts[1][len(parts[1])-4:] == ".txt" && len(parts[1]) > 4 {
				if params.reverse {
					return Error("Reverse défini deux fois")
				}
				params.reverse = true
				params.reverseFileName = parts[1]
			} else {
				return Error("Nom de fichier incorrect: " + parts[1])
			}
		}
	}

	if params.fontStyle && params.reverse { //Traitement des conflits des paramètres
		return Error("FS ne marche pas avec Reverse")
	}
	if params.output {
		if params.useColor || params.reverse || params.justify {
			return Error("Output ne marche qu'avec FS")
		}
	}
	return true
}

func TestColorValues(s string) bool {
	colors := []string{
		"red",
		"blue",
		"green",
		"yellow",
		"purple",
		"cyan",
		"white",
	}
	for _, val := range colors {
		if s == val {
			return true
		}
	}
	return false
}
