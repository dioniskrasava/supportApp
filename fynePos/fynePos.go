package fynepos

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/widget"
)

/*
Наша задача тут отработать позиционирование
*/
func App() {
	a := app.New()
	w := a.NewWindow("Hello World")

	w.SetContent(widget.NewLabel("Hello World!"))
	w.ShowAndRun()
}
