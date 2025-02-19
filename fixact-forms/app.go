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

	cont := awesomeShit(WIDTH, lbl, ent, btn)
	cont2 := awesomeShit(WIDTH, lbl2, ent2, btn2)

	globcont := container.NewVBox(cont, cont2)

	w.SetContent(globcont)
	w.ShowAndRun()
}

func awesomeShit(w float32, objects ...fyne.CanvasObject) *fyne.Container {
	// 3 колонки (размеры(%) 50:40:10)
	// ratios - соотношения
	ratios := []float32{0, 50, 80}

	percent := w / 100

	positionsX := getVertLinePos(ratios, percent) // позиции вертикал линий
	widthsX := getColumnWidths(positionsX, w)

	cont := getContainer(positionsX, widthsX, objects...)
	return cont

}

// получить
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
func getContainer(positionsX []float32, widthsX []float32, objects ...fyne.CanvasObject) *fyne.Container {
	cont := container.NewWithoutLayout()

	for k, v := range objects {
		v.Move(fyne.NewPos(positionsX[k], 10))
		v.Resize(fyne.NewSize(widthsX[k]-5, v.MinSize().Height))
		cont.Add(v)
	}

	return cont
}
