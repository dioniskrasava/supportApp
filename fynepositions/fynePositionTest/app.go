package fynepositiontest

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

var (
	// нужно будет разделить на макро, микро  минералы
	mineralsMacroList   = []string{"K", "Ca", "Si", "Mg", "Na", "P", "Cl"}
	mineralsMacroListUM = []string{"mg", "mg", "mg", "mg", "mg", "mg", "mg"} // единицы измерения
	mineralsMicroList   = []string{"Fe", "I", "Co", "Mn", "Cu", "Mo", "Se", "F", "Cr", "Zn"}
	mineralsMicroListUM = []string{"mg", "mcg", "mcg", "mg", "mcg", "mcg", "mcg", "mcg", "mcg", "mg"}                              // единицы измерения
	vitaminsList        = []string{"A", "b-car", "B1", "B2", "B4 Холин", "B5", "B6", "B9", "B12", "C", "D", "E", "H", "K", "PP"}   // Список меток для создания
	vitaminsListUM      = []string{"mcg", "mg", "mg", "mg", "mg", "mg", "mg", "mcg", "mcg", "mg", "mcg", "mg", "mcg", "mcg", "mg"} // единицы измерения

	green = color.RGBA{R: 80, G: 255, B: 80, A: 255}
	red   = color.RGBA{R: 255, G: 80, B: 80, A: 255}
	blue  = color.RGBA{R: 80, G: 80, B: 255, A: 255}
)

func createFormColumn(list []string, listUM []string, colorUM color.Color) *widget.Form {
	// Заголовок
	// Список лэйблов
	// список подсказок

	form := widget.NewForm()

	for i, value := range list {
		entry := widget.NewEntry()
		labelUM := canvas.NewText(listUM[i], colorUM)

		entryWithLabelUM := container.NewHBox(entry, labelUM)

		item := widget.NewFormItem(value, entryWithLabelUM)

		form.AppendItem(item)
	}

	return form

}

func App() {
	a := app.New()
	w := a.NewWindow("Grid Layout")
	w.Resize(fyne.NewSize(300, 300))

	newForm3 := createFormColumn(mineralsMacroList, mineralsMacroListUM, green)
	newForm2 := createFormColumn(mineralsMicroList, mineralsMicroListUM, red)
	newForm1 := createFormColumn(vitaminsList, vitaminsListUM, blue)

	globalcont := container.NewHBox(container.NewVBox(newForm1), container.NewVBox(newForm2), container.NewVBox(newForm3))

	w.SetContent(globalcont)
	w.ShowAndRun()
}
