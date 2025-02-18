package fynepositiontest

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func App() {
	a := app.New()
	w := a.NewWindow("Grid Layout")
	w.Resize(fyne.NewSize(300, 300))

	w1 := widget.NewEntry()
	w2 := widget.NewEntry()
	w3 := widget.NewEntry()

	item1 := widget.NewFormItem("A", w1)
	item2 := widget.NewFormItem("B", w2)
	item3 := widget.NewFormItem("C", w3)

	form := widget.NewForm(item1, item2, item3)

	//form.OnSubmit = func() {}
	//form.OnCancel = func() {}
	//------------------------------------

	w4 := widget.NewEntry()
	w5 := widget.NewEntry()
	w6 := widget.NewEntry()

	item4 := widget.NewFormItem("D", w4)
	item5 := widget.NewFormItem("E", w5)
	item6 := widget.NewFormItem("U", w6)

	form2 := widget.NewForm(item4, item5, item6)

	//form2.OnSubmit = func() {}
	//form2.OnCancel = func() {}
	//-------------------------------------

	globalcont := container.NewHBox(container.NewVBox(form), container.NewVBox(form2))

	w.SetContent(globalcont)
	w.ShowAndRun()
}
