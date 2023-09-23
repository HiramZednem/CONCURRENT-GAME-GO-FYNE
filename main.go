package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"pelota/scenes"
)

func main(){
	myApp := app.New()
	// windows
	game := createSimpleWindow( myApp )
	menu := createSimpleWindow( myApp )
	
	
	menuScene := scenes.NewMenuScene( menu )
	menuScene.Show()

	mainMenuScene := scenes.NewMainMenuScene( game )
	mainMenuScene.Show()
	
	
	menu.Show()
	myApp.Run()
}

func createSimpleWindow(app fyne.App) fyne.Window{
	window := app.NewWindow("Tux Revenge")
	window.CenterOnScreen()
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(800, 600))
	return window
}