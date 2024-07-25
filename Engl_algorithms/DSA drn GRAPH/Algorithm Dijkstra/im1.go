package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
)

func main() {
	a := app.New()
	w := a.NewWindow("Images")
	img := canvas.NewImageFromImage(generateImage())
	w.SetContent(img)
	w.Resize(fyne.NewSize(640, 480))
	w.ShowAndRun()
}

