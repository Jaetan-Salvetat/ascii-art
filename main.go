package main

import (
	"os"

	"./Base"
)

func main() {
	if len(os.Args) < 2 {
		print("No Arguments.")
		return
	}
	args := os.Args[1:]
	Base.PrintASCII(args)
}
