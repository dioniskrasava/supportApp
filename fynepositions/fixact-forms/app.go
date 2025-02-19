package fixactforms

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

const (
	WIDTH  float32 = 300
	HEIGHT float32 = 250
)

func App() {
	a := app.New()
	w := a.NewWindow("test")
	w.Resize(fyne.NewSize(WIDTH, HEIGHT))
	w.SetFixedSize(true)

	lbl := widget.NewLabel("Begin time :")
	lbl.TextStyle = fyne.TextStyle{Bold: true}

	ent := widget.NewEntry()
	ent.Resize(fyne.NewSize(ent.MinSize().Width, ent.MinSize().Height))

	btn := widget.NewButton("Click", func() {})
	btn.Resize(fyne.NewSize(btn.MinSize().Width, btn.MinSize().Height))

	//----------------------------------------
	lbl2 := widget.NewLabel("end :")
	lbl2.TextStyle = fyne.TextStyle{Bold: true}

	ent2 := widget.NewEntry()

	btn2 := widget.NewButton("!", func() {})

	// высоты строк
	h1 := float32(10)
	h2 := float32(50)

	globContainer := container.NewWithoutLayout()

	awesomeShit(globContainer, WIDTH, h1, lbl, ent, btn)
	awesomeShit(globContainer, WIDTH, h2, lbl2, ent2, btn2)

	w.SetContent(globContainer)
	w.ShowAndRun()
}

// название классное
// принимает глобальный кастомный контейнер container.NewWithoutLayout(), ширину окна, высоту позиции строки виджетов (пока еще недоработан этот момент)
// и пачку канвас объектов которые нужно расстрочить
func awesomeShit(globContainer *fyne.Container, widthWindow float32, heightWidg float32, objects ...fyne.CanvasObject) {

	// ratios - процентные соотношения, на которые делится окно
	ratios := []float32{0, 50, 80}

	// высчитываем 1 процент
	percent := widthWindow / 100

	// получаем позиции вертик линий
	positionsX := getVertLinePos(ratios, percent)
	// получаем конкретные координаты вертикальных линий по Х
	widthsX := getColumnWidths(positionsX, widthWindow)
	// стругаем контейнер с уже позиционированными виджетами
	addGlobContainer(globContainer, positionsX, widthsX, heightWidg, objects...)

}

// получить позиции вертикальных границ
// вспомогательная функция для awesomeShit
func getVertLinePos(ratios []float32, percent float32) []float32 {
	positionsX := []float32{} // 0, 250, 400
	// вычисляем позиции вертикальных рамок
	for _, v := range ratios {
		posX := percent * v
		positionsX = append(positionsX, posX)
	}

	fmt.Println("Позиции вертикальных линий : ", positionsX)
	return positionsX
}

// получить позиции по оси Х
// вспомогательная функция для awesomeShit
func getColumnWidths(positionsX []float32, w float32) []float32 {
	widthsX := []float32{}
	// вычисляем ширины колонок
	for k, _ := range positionsX {
		// если последний элемент
		if k == len(positionsX)-1 {
			width := w - positionsX[k]
			widthsX = append(widthsX, width)
		} else {
			width := positionsX[k+1] - positionsX[k]
			widthsX = append(widthsX, width)
		}
	}
	fmt.Println("Ширины колонок : ", widthsX)
	return widthsX
}

// добавление объектов на кастомную сетку
// вспомогательная функция для awesomeShit
func addGlobContainer(globContainer *fyne.Container, positionsX []float32, widthsX []float32, heightWidg float32, objects ...fyne.CanvasObject) {

	for k, v := range objects {
		v.Move(fyne.NewPos(positionsX[k], heightWidg))
		v.Resize(fyne.NewSize(widthsX[k]-5, v.MinSize().Height))
		globContainer.Add(v)
	}
}
