package fixact

import (
	"database/sql"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	_ "github.com/mattn/go-sqlite3"
)

type Activity struct {
	Type      string
	StartTime string
	EndTime   string
	TotalTime string
	Comment   string
}

func App() {
	// Инициализация приложения Fyne
	myApp := app.New()
	myWindow := myApp.NewWindow("Activity Tracker")

	// Подключение к базе данных SQLite
	db, err := sql.Open("sqlite3", "./fixact/activities.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Создание таблицы, если она не существует
	createTableSQL := `CREATE TABLE IF NOT EXISTS activities (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	type TEXT,
	start_time TEXT,
	end_time TEXT,
	total_time TEXT,
	comment TEXT
);`
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}

	// Элементы интерфейса
	activityType := widget.NewSelect([]string{"Книга", "Код", "Видео"}, func(value string) {})
	startTime := widget.NewEntry()
	endTime := widget.NewEntry()
	totalTime := widget.NewEntry()
	comment := widget.NewMultiLineEntry()
	addButton := widget.NewButton("Добавить активность", func() {
		activity := Activity{
			Type:      activityType.Selected,
			StartTime: startTime.Text,
			EndTime:   endTime.Text,
			TotalTime: totalTime.Text,
			Comment:   comment.Text,
		}

		// Вставка данных в базу данных
		insertSQL := `INSERT INTO activities (type, start_time, end_time, total_time, comment) VALUES (?, ?, ?, ?, ?)`
		_, err := db.Exec(insertSQL, activity.Type, activity.StartTime, activity.EndTime, activity.TotalTime, activity.Comment)
		if err != nil {
			log.Fatal(err)
		}

		// Очистка полей после добавления
		activityType.SetSelected("")
		startTime.SetText("")
		endTime.SetText("")
		totalTime.SetText("")
		comment.SetText("")

		fmt.Println("Активность добавлена!")
	})

	// Создание контейнера с элементами интерфейса
	content := container.NewVBox(
		widget.NewLabel("Тип активности:"),
		activityType,
		widget.NewLabel("Время начала:"),
		startTime,
		widget.NewLabel("Время окончания:"),
		endTime,
		widget.NewLabel("Общее время:"),
		totalTime,
		widget.NewLabel("Комментарий:"),
		comment,
		addButton,
	)

	// Установка контента в окно
	myWindow.SetContent(content)
	myWindow.Resize(fyne.NewSize(400, 300))
	myWindow.ShowAndRun()
}
