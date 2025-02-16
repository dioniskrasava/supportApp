package fynepos

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func App() {
	a := app.New()
	w := a.NewWindow("apppppplication")
	w.Resize(fyne.NewSize(500, 500))

	data := []string{"A", "B", "C", "Холин"}

	table := widget.NewTable(
		func() (int, int) {
			return len(data), 1
		},

		func() fyne.CanvasObject {
			return widget.NewLabel("Default text")
		},

		func(tci widget.TableCellID, co fyne.CanvasObject) {
			co.(*widget.Label).SetText(data[tci.Row])
		},
	)

	table2 := widget.NewTable(
		func() (int, int) {
			return len(data), 1
		},

		func() fyne.CanvasObject {
			entry := widget.NewEntry()
			entry.SetPlaceHolder("0")
			// Устанавливаем минимальную ширину для Entry
			return container.NewHBox(container.NewGridWrap(fyne.NewSize(75, entry.MinSize().Height), entry))
		},

		func(tci widget.TableCellID, co fyne.CanvasObject) {
			//co.(*widget.Label).SetText(data[tci.Row])
		},
	)

	table3 := widget.NewTable(
		func() (int, int) {
			return len(data), 1
		},

		func() fyne.CanvasObject {
			return widget.NewLabel("mg")
		},

		func(tci widget.TableCellID, co fyne.CanvasObject) {

		},
	)

	cont := container.NewAdaptiveGrid(3, table, table2, table3)

	w.SetContent(cont)

	widthOneTable := table.Size()
	fmt.Println(widthOneTable)

	w.Show()
	a.Run()
}
