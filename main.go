package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"pelota/scenes"
)

func main(){
	myApp := app.New()
	myWindow := myApp.NewWindow("Juego Pelota")
	myWindow.CenterOnScreen()
	myWindow.SetFixedSize(true)
	myWindow.Resize(fyne.NewSize(800, 600))

	//Cargar y mostrar la escena principal
	mainMenuScene := scenes.NewMainMenuScene(myWindow)
	mainMenuScene.Show()
	myWindow.ShowAndRun()
}
