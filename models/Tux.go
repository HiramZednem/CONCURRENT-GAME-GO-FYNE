package models

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Tux struct {
	posX, posY, direction float32
	pause bool	
	pel *canvas.Image
}

func NewTux(posx float32, posy float32, img *canvas.Image) *Tux {
	return &Tux{
		posX: posx,
		posY: posy,
		pause: false,
		pel: img,
		direction: 0,
	}
}

func (t *Tux) Run() {
	var incX float32 = 50
	for true { // TODO: isColisioned
		for !t.pause {
			/*
			Aqui hay que cambiar la logica, en la scene agregare dos botones, < y >
			el valor de tux se va a multiplicar por "direction", que va a ser un valor
			entre -1 y 1, asi logrando que se mueva entre derecha e izquierda
			
			Ahora, quiero hacer que si la pos.x llega a un limite, como 50 a 700, entonces sume 0
			para que no se mueva hasta que pulses la direccion hacia el otro lado
			*/


			if t.posX < 50 || t.posX > 650 {
				incX *= -1	
			}
			t.posX += incX
			fmt.Println(t.posX, t.posY)
			t.pel.Move(fyne.NewPos(t.posX,t.posY))
			time.Sleep(100 * time.Millisecond)
		}
	}
}

func (t *Tux) SetPause(pause bool) {
	t.pause = pause
}
func (t *Tux) GetPause() bool {
	return t.pause
}