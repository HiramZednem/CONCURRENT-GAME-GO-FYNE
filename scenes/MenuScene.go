package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type MenuScene struct {
	window fyne.Window
}

func NewMenuScene(window fyne.Window) *MenuScene {
	return &MenuScene{window: window,}
}

func (menuScene *MenuScene) Show() {
	backgroundImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/background.jpg") )
    backgroundImage.Resize(fyne.NewSize(800,600))
	backgroundImage.Move( fyne.NewPos(0,0) )
	
	botonIniciar := widget.NewButton("Start Game", menuScene.StartGame)
	botonIniciar.Resize(fyne.NewSize(150,30))
	botonIniciar.Move(fyne.NewPos(300,10))

	menuScene.window.SetContent(container.NewWithoutLayout(backgroundImage, botonIniciar)) 
}

func (menuScene *MenuScene) StartGame() {
	go p.Run()
}

