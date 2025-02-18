package fynepositiontest

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func App() {
	a := app.New()
	w := a.NewWindow("Grid Layout")

	// Создаем виджеты
	button1 := widget.NewButton("Button 1", nil)
	button2 := widget.NewButton("Button 2", nil)
	button3 := widget.NewButton("Button 3", nil)

	// Создаем сетку с тремя столбцами
	grid := layout.NewGridLayout(3) // 3 столбца

	// Устанавливаем ширину столбцов
	grid.SetColumnWidths([]int{2, 2, 1}) // Ширина столбцов: 2, 2, 1

	// Добавляем виджеты в сетку
	content := container.New(grid, button1, button2, button3)

	w.SetContent(content)
	w.ShowAndRun()
}
