/*
	Inicializo un timer
*/

package scenes

import (
	"TuxGame/driver"
	"TuxGame/models"
	"time"

	// "time"

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
var w *models.Windows
var c *driver.CollisionDriver

func NewGameScene(window fyne.Window) *GameScene {
	scene := &GameScene{window: window}
    scene.Render()
	scene.StartGame()
    return scene
}

func (s *GameScene) Render() {
	backgroundImage := canvas.NewImageFromURI( storage.NewFileURI("./assets/darkPurple.png") )
    backgroundImage.Resize(fyne.NewSize(800,600))
	backgroundImage.Move( fyne.NewPos(0,0) )
	// Models Render
	tuxPeel := createPeel("./assets/tux.png", 100, 100, 100, 450)
	t = models.NewTux( 350, 450, tuxPeel )

	windowsPeel := createPeel("./assets/windows.png", 100, 100, 100, 50)
	w = models.NewWindows( 350, 600, windowsPeel, t )

	//CollisionDriver
	c = driver.NewCollisionDriver(t, w)

	// Buttons Render
	btnLeft := widget.NewButton("<", t.GoLeft)
	btnLeft.Resize(fyne.NewSize(50,50))
	btnLeft.Move(fyne.NewPos(350,550))

	btnRigth := widget.NewButton(">", t.GoRigth)
	btnRigth.Resize(fyne.NewSize(50,50))
	btnRigth.Move(fyne.NewPos(400,550))

	s.window.SetContent(container.NewWithoutLayout(backgroundImage, tuxPeel, windowsPeel, btnLeft, btnRigth)) 
}

func (s *GameScene) StartGame() {
	go t.Run()
	go w.Run()
	go c.Run()
	go s.checkGameOver() // Esto se podia hacer utilizando un channel, pero como no lo hemos visto no lo implemente
}

func (s *GameScene) StopGame() {
	t.SetRunning(!t.GetRunning())
	w.SetRunning(!w.GetRunning())
}

func (s *GameScene) checkGameOver() {
	running := true
	for running {
		if(c.GetGameOver()) {
			running = false
			time.Sleep(1000 * time.Millisecond)
			NewGameOverScene(s.window)
			
		}
	}
}

func createPeel( fileUri string, sizeX float32, sizeY float32, posX float32, posY float32 ) *canvas.Image {
	image := canvas.NewImageFromURI(storage.NewFileURI(fileUri))
	image.Resize(fyne.NewSize(sizeX,sizeY))
	image.Move(fyne.NewPos(posX,posY))
	return image
}