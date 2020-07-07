package main

import (
	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	)

func main()  {
	a := app.New()
	w:= a.NewWindow("Hello")
	w.SetContent(widget.NewVBox(
		widget.NewLabel("hello k6"),
		widget.NewButton("Quit", func() {
			a.Quit()
		})))
	w.ShowAndRun()

}



