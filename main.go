package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/widgets"
)

func main() {
	// init app
	app := createApp()
	mainWindow := createWindow("Save", 400, 400)
	layout := createLayout()
	mainWidget := createWidget()
	mainWidget.SetLayout(layout)

	// create component
	input := widgets.NewQLineEdit(nil)
	input.SetPlaceholderText("1. write something")
	layout.AddWidget(input, 0, 0)

	filedl := widgets.NewQFileDialog(nil, core.Qt__Dialog)

	button := widgets.NewQPushButton2("2. click me", nil)
	layout.AddWidget(button, 0, 0)

	button.ConnectClicked(func(checked bool) {
		//filedl.Exec()
		var str string
		for _, s := range filedl.GetOpenFileNames(nil, "asdfg", "", "", "", widgets.QFileDialog__ReadOnly) {
			str += (s + "\n")
		}
		widgets.QMessageBox_Information(nil, "OK", str, widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)
		//widgets.QMessageBox_Information(nil, "OK", input.Text(), widgets.QMessageBox__Ok, widgets.QMessageBox__Ok)

	})

	// show app
	mainWindow.SetCentralWidget(mainWidget)
	mainWindow.Show()
	app.Exec()
}
