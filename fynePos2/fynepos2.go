package fynepos2

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
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
)

func App() {

	globalListLabel := createGlobalList(vitaminsList, mineralsMacroList, mineralsMicroList)
	globalListLabelUM := createGlobalList(vitaminsListUM, mineralsMacroListUM, mineralsMicroListUM)

	createNewCont(globalListLabel, globalListLabelUM)

}

func createGlobalList(list1 []string, list2 []string, list3 []string) []string {
	// Объединяем все три среза в один
	result := append(list1, list2...)
	result = append(result, list3...)
	return result
}

func createNewCont(listLabel []string, listUM []string) {
	// получаем все списки элементов + списки наименований

	//vit := "Витамины"
	//macr := "Макроэлементы"
	//micr := "Микроэлементы"

	column1 := container.NewVBox() // Контейнер для всех строк
	column2 := container.NewVBox() // Контейнер для всех строк
	column3 := container.NewVBox() // Контейнер для всех строк
	column4 := container.NewVBox() // Контейнер для всех строк

	columns := []*fyne.Container{column1, column2, column3, column4}

	n := 0 // вспомогательная переменная для того, что знать какой лэйб вытаскивать из общего списка лейблов
	// 4 столба по 9 строк
	for i := 0; i < 4; i++ {
		for j := 0; j < 9; j++ {

			if i == 0 && j == 0 {
				//vitamins
				label := widget.NewLabel("Витамины")
				columns[i].Add(label)
			} else if i == 1 && j == 7 {
				//macro
				label := widget.NewLabel("Макроэлементы")
				columns[i].Add(label)
			} else if i == 2 && j == 6 {
				//micro
				label := widget.NewLabel("Микроэлементы")
				columns[i].Add(label)
			} else if i == 3 && j == 8 {
				// пустое место (элементы кончились)
				break
			}
			label := widget.NewLabel(listLabel[n])
			entry := widget.NewEntry()
			labelUM := widget.NewLabel(labelsUM[i]) // Unit Measure - единицы измерения
			n++

			// Сохраняем ссылку на поле ввода в карту
			entries[labelText] = entry

			//row := container.NewGridWithColumns(4, label, layout.NewSpacer(), entry, labelUM)
			row := container.NewHBox(label, layout.NewSpacer(), entry, labelUM)
			rows.Add(row)
			rows.Add(widget.NewSeparator())

			// создание столбца
			label := widget.NewLabel("")
			entry := widget.NewEntry()
		}
		fmt.Println()
	}
}
