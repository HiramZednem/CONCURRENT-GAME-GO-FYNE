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
	gameOverImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/gameOver.png") )
    gameOverImage.Resize(fyne.NewSize(200,150))
	gameOverImage.Move( fyne.NewPos(300,225) )

	backgroundImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/darkPurple.png") )
    backgroundImage.Resize(fyne.NewSize(800,600))
	backgroundImage.Move( fyne.NewPos(0,0) )
	
	btnGoMenu := widget.NewButton("Go Menu", s.goMenu)
	btnGoMenu.Resize(fyne.NewSize(150,30))
	btnGoMenu.Move(fyne.NewPos(325, 50))

	btnRestart := widget.NewButton("Restart Game", s.restart)
	btnRestart.Resize(fyne.NewSize(150,30))
	btnRestart.Move(fyne.NewPos(325, 10))
	

	s.window.SetContent(container.NewWithoutLayout(backgroundImage, gameOverImage,btnGoMenu, btnRestart)) 
}

func (s *GameOverScene) goMenu() {
	NewMenuScene(s.window)
}
func (s *GameOverScene) restart() {
	NewGameScene(s.window)
}