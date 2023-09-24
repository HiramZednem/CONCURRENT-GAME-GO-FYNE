package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"pelota/scenes"
)

func main(){
	myApp := app.New()
	menu := createSimpleWindow( myApp )
	scenes.NewMenuScene( menu, createSimpleWindow( myApp ) )
	menu.ShowAndRun()
}

func createSimpleWindow(app fyne.App) fyne.Window{
	window := app.NewWindow("Tux Revenge")
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(800, 600))
	return window
}