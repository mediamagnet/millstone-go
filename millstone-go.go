package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"github.com/c-bata/go-prompt"
	"io"
	"millstone-go/lib"
)

var parseWarnC int
var parseErrC int
var parseWarnD []string
var parseErrD []string

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "filepath", Description: "What log file do you want to parse?"},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func main() {
	var file string
	a := app.New()
	w := a.NewWindow("Millstone-Go")
	w.Resize(fyne.NewSize(600, 400))
	btn := widget.NewButton("Open log file", func() {
		file_Dialog := dialog.NewFileOpen(
			func(r fyne.URIReadCloser, _ error) {
				file, _ := io.ReadAll(r)
				result := fyne.NewStaticResource("name", file)
				entry := widget.NewMultiLineEntry()
				entry.SetText(string(result.StaticContent))
				w := fyne.CurrentApp().NewWindow(
					string(result.StaticName))
				w.SetContent(container.NewScroll(entry))
				w.Resize(fyne.NewSize(600, 400))
				w.Show()
			}, w)
		file_Dialog.SetFilter(
			storage.NewExtensionFileFilter([]string{".txt"}))
		file_Dialog.Show()
	})
	w.SetContent(container.NewVBox(btn))
	w.ShowAndRun()
	parseWarnC, parseErrC, parseWarnD, parseErrD, err := lib.LogParse(file)
	if err != nil {
		return
	}

	fmt.Printf("%v, %v, %v, %v", parseWarnC, parseErrC, parseWarnD, parseErrD)
}
