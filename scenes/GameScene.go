/*
	Inicializo un timer
*/

package scenes

import (
	"TuxGame/models"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
)

type GameScene struct {
	window fyne.Window
}

var t *models.Tux

func NewGameScene(window fyne.Window) *GameScene {
	scene := &GameScene{window: window}
    scene.Render()
	scene.StartGame()
    return scene
}

func (s *GameScene) Render() {
	tuxPeel := createPeel("./assets/tux.png", 100, 100, 100, 500)
	t = models.NewTux(350,450,tuxPeel)
	
	botonDetener := widget.NewButton("||", s.StopGame)
	botonDetener.Resize(fyne.NewSize(50,50))
	botonDetener.Move(fyne.NewPos(750,550))

	btnLeft := widget.NewButton("<", t.GoLeft)
	btnLeft.Resize(fyne.NewSize(50,50))
	btnLeft.Move(fyne.NewPos(350,550))

	btnRigth := widget.NewButton(">", t.GoRigth)
	btnRigth.Resize(fyne.NewSize(50,50))
	btnRigth.Move(fyne.NewPos(400,550))

	s.window.SetContent(container.NewWithoutLayout(tuxPeel, botonDetener, btnLeft, btnRigth)) 
}

func (s *GameScene) StartGame() {
	go t.Run()
	// TODO: Crear modelo windows, y cada vez que llegue a posY 450 (osea toque el bottom), regrese a posY 0, y tenga su movimiento de bajar siempre comprobando
	// que window.collisioned, si colisiona con tux, el juego acaba
	// TODO: Crear Timer, para ver cuanto tiempo sobrevives y cuando window.collisioned, te diga has sobrevivido 1:30m
}

func (s *GameScene) StopGame() {
	t.SetRunning(!t.GetRunning())
}

func createPeel( fileUri string, sizeX float32, sizeY float32, posX float32, posY float32 ) *canvas.Image {
	image := canvas.NewImageFromURI(storage.NewFileURI(fileUri))
	image.Resize(fyne.NewSize(sizeX,sizeY))
	image.Move(fyne.NewPos(posX,posY))
	return image
}