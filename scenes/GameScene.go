/*
	Inicializo un timer
*/

package scenes

import (
	"pelota/models"

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
	t = models.NewTux(100,500,tuxPeel)
	
	botonDetener := widget.NewButton("Stop Game", s.StopGame)
	botonDetener.Resize(fyne.NewSize(150,30))
	botonDetener.Move(fyne.NewPos(300,50))


	s.window.SetContent(container.NewWithoutLayout(tuxPeel, botonDetener)) 
}

func (s *GameScene) StartGame() {
	go t.Run()
	// TODO: Crear modelo windows, y cada vez que llegue a posY 800 (osea toque el bottom), regrese a posY 0, y tenga su movimiento de bajar siempre comprobando
	// que window.collisioned, si colisiona con tux, el juego acaba
	// TODO: Crear Timer, para ver cuanto tiempo sobrevives y cuando window.collisioned, te diga has sobrevivido 1:30m
	// Ahora, para que no se vuelva repetitivo que voy a hacer?, no importa la jugabilidad xd, asi que a csm
}

func (s *GameScene) StopGame() {
	t.SetPause(!t.GetPause())
}

func createPeel( fileUri string, sizeX float32, sizeY float32, posX float32, posY float32 ) *canvas.Image {
	image := canvas.NewImageFromURI(storage.NewFileURI(fileUri))
	image.Resize(fyne.NewSize(sizeX,sizeY))
	image.Move(fyne.NewPos(posX,posY))
	return image
}