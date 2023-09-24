package models

import (
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Tux struct {
	posX, posY, direction float32
	running bool
	pel *canvas.Image
}

func NewTux(posx float32, posy float32, img *canvas.Image) *Tux {
	return &Tux{
		posX: posx,
		posY: posy,
		running: true,
		pel: img,
		direction: 0,
	}
}

func (t *Tux) GoRigth() {
	t.direction = 1
}

func (t *Tux) GoLeft () {
	t.direction = -1
}

func (t *Tux) Run() {
	for true { // TODO: isColisioned
		for t.running {
			var incX float32 = 50
			incX *= t.direction

			if t.posX < 50 {
				t.posX = 50
			} else if t.posX > 650 {
				t.posX = 650
			}

			t.posX += incX
			t.pel.Move(fyne.NewPos(t.posX,t.posY))
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (t *Tux) SetRunning(pause bool) {
	t.running = pause
}
func (t *Tux) GetRunning() bool {
	return t.running
}
func (t *Tux) GetPosY() float32 {
	return t.posY
}
func (t *Tux) GetPosX() float32 {
	return t.posX
}