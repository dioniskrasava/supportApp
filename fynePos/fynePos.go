package fynepos

import (
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	WIDTH_WINDOW  = 400
	HEIGHT_WINDOW = 600
)

var (
	vitaminsList   = []string{"A", "b-car", "B1", "B2", "B4 Холин", "B5", "B6", "B9", "B12", "C", "D", "E", "H", "K", "PP"}   // Список меток для создания
	vitaminsListUM = []string{"mcg", "mg", "mg", "mg", "mg", "mg", "mg", "mcg", "mcg", "mg", "mcg", "mg", "mcg", "mcg", "mg"} // единицы измерен
)

func App() {
	a := app.New()
	w := a.NewWindow("Table with Labels and Entries")
	w.Resize(fyne.NewSize(WIDTH_WINDOW, HEIGHT_WINDOW))

	// Запрещаем изменение размера окна
	w.SetFixedSize(true)

	table, entries := createTable(vitaminsList, vitaminsListUM)

	for key, value := range entries {
		fmt.Println(key, value.Text)
		value.OnChanged = func(s string) {
			log.Printf("%s changed to %s\n", key, s)
		}
	}

	w.SetContent(table)
	w.Show()
	a.Run()
}

func createTable(labels []string, labelsUM []string) (*widget.Table, map[string]*widget.Entry) {

	entries := make(map[string]*widget.Entry) // Карта для хранения полей ввода

	// Создаем таблицу с двумя столбцами
	table := widget.NewTable(
		func() (int, int) {
			return len(labels), 3 // 2 столбца: один для меток, один для полей ввода
		},
		func() fyne.CanvasObject {
			// Возвращаем контейнер, который может содержать либо Label, либо Entry
			return container.NewStack(widget.NewLabel(""), widget.NewEntry())
		},
		func(tci widget.TableCellID, co fyne.CanvasObject) {
			// Настройка содержимого ячеек
			stack := co.(*fyne.Container)
			if tci.Col == 0 {
				// Первый столбец: метки
				label := stack.Objects[0].(*widget.Label)
				label.SetText(labels[tci.Row])
				label.Show()
				stack.Objects[1].(*widget.Entry).Hide()
			} else if tci.Col == 1 {
				// Второй столбец: поля ввода
				entry := stack.Objects[1].(*widget.Entry)
				entry.SetPlaceHolder("0")
				entry.Show()
				stack.Objects[0].(*widget.Label).Hide()
				entries[labels[tci.Row]] = entry // ДОБАВЛЯЕМ НАШИ ЕНТРИ ПО КЛЮЧУ (ЗНАЧЕНИЕ ЛЕЙБЛА)
			} else if tci.Col == 2 {
				label := stack.Objects[0].(*widget.Label)
				label.SetText(labelsUM[tci.Row])
				label.Show()
				stack.Objects[1].(*widget.Entry).Hide()
			}
		},
	)

	// Устанавливаем ширину столбцов
	table.SetColumnWidth(0, 75) // Первый столбец шире (200 пикселей)
	table.SetColumnWidth(1, 50) // Второй столбец уже (100 пикселей)

	log.Println(entries)

	return table, entries
}
