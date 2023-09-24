package scenes

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type GameOverScene struct {
	window fyne.Window
}

func NewGameOverScene(fyneWindow fyne.Window) *GameOverScene {
	scene := &GameOverScene{window: fyneWindow}
	scene.Render()
	return scene
}

func (s *GameOverScene) Render() {
	backgroundImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/gameOver.png") )
    backgroundImage.Resize(fyne.NewSize(800,600))
	backgroundImage.Move( fyne.NewPos(0,0) )
	
	btnGoMenu := widget.NewButton("Go Menu", s.goMenu)
	btnGoMenu.Resize(fyne.NewSize(150,30))
	btnGoMenu.Move(fyne.NewPos(300,10))

	s.window.SetContent(container.NewWithoutLayout(backgroundImage, btnGoMenu)) 
}

func (s *GameOverScene) goMenu() {
	NewMenuScene(s.window)
}