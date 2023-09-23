package models

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
)

type Pelota struct {
	posX, posY float32
	status bool	
	pel *canvas.Image
}

func NewPelota(posx float32, posy float32, img *canvas.Image) *Pelota {
	return &Pelota{
		posX: posx,
		posY: posy,
		status: true,
		pel: img,
	}
}

func (p *Pelota) Run() {
	var incX float32 = 50
	// var incY float32 = 50
	p.status = true
	for p.status {
		if p.posX < 50 || p.posX >740 {
			incX *= -1	
		}
		// if p.posY < 50 || p.posY > 540 {
		// 	incY *= -1
		// }
		p.posX += incX
		// p.posY += incY
		fmt.Println(p.posX, p.posY)
		p.pel.Move(fyne.NewPos(p.posX,p.posY))
		time.Sleep(100 * time.Millisecond)
	}	
}
func (p *Pelota) SetStatus(status bool) {
	p.status = status
}