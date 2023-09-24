package models

import (
	"math/rand"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

var t *Tux

type Windows struct {
	posX, posY, direction float32
	running bool	
	pel *canvas.Image
}

func NewWindows(posx float32, posy float32, img *canvas.Image, tux *Tux) *Windows {
	t = tux
	return &Windows{
		posX: posx,
		posY: posy,
		running: true,
		pel: img,
	}
}

func (w *Windows) Run() {
	for w.running {
		var inc float32 = 50

		if w.posY > 500 {
			w.posY = -50
			w.posX = float32((rand.Intn(12) + 1) * 50)
		}
		
		w.posY += inc
		w.pel.Move(fyne.NewPos(w.posX,w.posY))
		time.Sleep(100 * time.Millisecond)
	}
}

func (w *Windows) SetRunning(pause bool) {
	w.running = pause
}
func (w *Windows) GetRunning() bool {
	return w.running
}
func (w *Windows) GetPosY() float32 {
	return w.posY
}
func (w *Windows) GetPosX() float32 {
	return w.posX
}

