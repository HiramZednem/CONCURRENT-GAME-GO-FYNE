package scenes

import (
	_ "fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	"pelota/models"
)

type GameScene struct {
	window fyne.Window
}

var p *models.Pelota

func NewGameScene(window fyne.Window) *GameScene {
	scene := &GameScene{window: window}
    scene.Render()
	scene.StartGame()
    return scene
}

func (s *GameScene) Render() {

	pelota := createPeel("./assets/tux.png", 100, 100, 100, 500)
	
	//Creamos el modelo
	p = models.NewPelota(100,500,pelota)
	
	botonIniciar := widget.NewButton("Start Game", s.StartGame)
	botonIniciar.Resize(fyne.NewSize(150,30))
	botonIniciar.Move(fyne.NewPos(300,10))

	botonDetener := widget.NewButton("Stop Game", s.StopGame)
	botonDetener.Resize(fyne.NewSize(150,30))
	botonDetener.Move(fyne.NewPos(300,50))


	s.window.SetContent(container.NewWithoutLayout(pelota, botonIniciar, botonDetener)) 
}

func (s *GameScene) StartGame() {
	go p.Run()
}

func (s *GameScene) StopGame() {
	p.SetStatus(false)
}

func createPeel( fileUri string, sizeX float32, sizeY float32, posX float32, posY float32 ) *canvas.Image {
	image := canvas.NewImageFromURI(storage.NewFileURI(fileUri))
	image.Resize(fyne.NewSize(sizeX,sizeY))
	image.Move(fyne.NewPos(posX,posY))
	return image
}