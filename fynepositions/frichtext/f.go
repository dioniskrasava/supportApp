package frichtext

import (
	"image/color"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func App() {
	a := app.New()
	w := a.NewWindow("CanvasText с кастомным цветом")

	text := canvas.NewText("Привет, Fyne!", color.RGBA{R: 0, G: 128, B: 0, A: 255}) // Зеленый цвет
	text.TextSize = 24                                                              // Размер шрифта

	w.SetContent(container.NewCenter(text))
	w.ShowAndRun()
}
