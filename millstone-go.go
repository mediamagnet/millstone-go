package main

import (
	"github.com/c-bata/go-prompt"
	"github.com/fatih/color"
	"millstone-go/lib"
)

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "filepath", Description: "What log file do you want to parse?"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	ask := color.New(color.FgCyan, color.Bold).PrintlnFunc()
	ask("What log file do you want to open?")
	file := prompt.Input("> ", completer)
	lib.LogParse(file)
}
