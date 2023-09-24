package driver

import "TuxGame/models"

var t *models.Tux
var w *models.Windows

type CollisionDriver struct { }

func NewCollisionDriver(tux *models.Tux, windows *models.Windows) *CollisionDriver {
	t = tux
	w = windows
	return &CollisionDriver{}
}

func (c *CollisionDriver) Run() {
	for {
		if w.GetPosY() >= 400 {
			if w.GetPosX() >= t.GetPosX()-50 && w.GetPosX() <= t.GetPosX()+50 {
				w.SetRunning(false)
				t.SetRunning(false)
			}
		} 
	}
}