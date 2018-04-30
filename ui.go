package main

import (
	"github.com/therecipe/qt/widgets"
	"os"
)

func createApp() *widgets.QApplication {
	return widgets.NewQApplication(len(os.Args), os.Args)
}

func createWindow(title string, xSize, ySize int) *widgets.QMainWindow {
	window := widgets.NewQMainWindow(nil, 0)
	window.SetWindowTitle(title)
	window.SetMinimumSize2(xSize, ySize)
	return window
}

func createLayout() *widgets.QVBoxLayout {
	return widgets.NewQVBoxLayout()
}

func createWidget() *widgets.QWidget {
	return widgets.NewQWidget(nil, 0)
}