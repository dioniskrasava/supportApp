package fynepos

/*
в общем это приложение преследовало цель создать отображение страницы добавления нового
продукта в моем приложении о питании в виде таблицы

но проблема в том, что виджет таблицы работает уж очень специфично
что в конечном итоге приводит к интересным результатам.

Например при добавлении значения в ентри и скроллинге, данные из ентри могут исчезнуть и не зафиксироваться.
Видимо это связано с тем, что таблица постоянно обновлятеся. По крайней мере так видно по логам.

Поэтому видимо отображение в виде таблицы можно считать неуспешным
*/

import (
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	WIDTH_WINDOW  = 400
	HEIGHT_WINDOW = 400
)

var (
	vitaminsList   = []string{"A", "b-car", "B1", "B2", "B4 Холин", "B5", "B6", "B9", "B12", "C", "D", "E", "H", "K", "PP"}   // Список меток для создания
	vitaminsListUM = []string{"mcg", "mg", "mg", "mg", "mg", "mg", "mg", "mcg", "mcg", "mg", "mcg", "mg", "mcg", "mcg", "mg"} // единицы измерен

	entriesVitamins = make([]*widget.Entry, len(vitaminsList)) // создаем срез указателей на *widget.Entry длиной равной длине списка витаминов

)

func App() {
	a := app.New()
	w := a.NewWindow("Table with Labels and Entries")
	w.Resize(fyne.NewSize(WIDTH_WINDOW, HEIGHT_WINDOW))

	w.SetFixedSize(true)

	table := createTable(vitaminsList, vitaminsListUM)

	// Добавляем кнопку для проверки значений
	checkButton := widget.NewButton("Check Values", func() {
		for i, entry := range entriesVitamins {
			if entry != nil {
				log.Printf("Vitamin %s: %s\n", vitaminsList[i], entry.Text)
			} else {
				log.Printf("Entry for vitamin %s is nil\n", vitaminsList[i])
			}
		}
	})

	// Используем container.NewBorder для размещения таблицы и кнопки
	content := container.NewBorder(
		nil,         // Верхний элемент (nil, так как ничего не нужно)
		checkButton, // Нижний элемент (кнопка)
		nil,         // Левый элемент (nil)
		nil,         // Правый элемент (nil)
		table,       // Центральный элемент (таблица)
	)

	w.SetContent(content)
	w.Show()
	a.Run()
}

func createTable(labels []string, labelsUM []string) *widget.Table {
	table := widget.NewTable(
		func() (int, int) {
			return len(labels), 3 // 2 столбца: один для меток, один для полей ввода
		},
		func() fyne.CanvasObject {
			return container.NewStack(widget.NewLabel(""), widget.NewEntry())
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			stack := co.(*fyne.Container)
			if tci.Col == 0 {
				label := stack.Objects[0].(*widget.Label)
				label.SetText(labels[tci.Row])
				label.Show()
				stack.Objects[1].(*widget.Entry).Hide()
			} else if tci.Col == 1 {
				entry := stack.Objects[1].(*widget.Entry)
				entry.SetPlaceHolder("0")
				entry.Show()
				stack.Objects[0].(*widget.Label).Hide()

				entriesVitamins[tci.Row] = entry
				log.Printf("Entry added to entriesVitamins at index %d\n", tci.Row)
			} else if tci.Col == 2 {
				label := stack.Objects[0].(*widget.Label)
				label.SetText(labelsUM[tci.Row])
				label.Show()
				stack.Objects[1].(*widget.Entry).Hide()
			}
		},
	)

	table.SetColumnWidth(0, 75)
	table.SetColumnWidth(1, 50)
	table.SetColumnWidth(2, 50)

	return table
}
