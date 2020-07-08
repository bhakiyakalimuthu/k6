package main

import (
	"fyne.io/fyne"
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	)

func main()  {
	a := app.New()
	w:= a.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabelWithStyle("hello k6",2,fyne.TextStyle{
			Bold:      true,
			Italic:    false,
			Monospace: true,
		}),
		widget.NewButton("Quit", func() {
			a.Quit()
		})))
	w.ShowAndRun()

}



