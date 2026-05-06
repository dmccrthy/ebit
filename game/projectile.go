package game

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Projectile struct {
	X      float64
	Y      float64
	Xdirection float64
	Ydirection float64
	Width  int
	Height int
	Image  *ebiten.Image
}

func NewProjectile(x, y, xd, yd float64, image *ebiten.Image) *Projectile {
	return &Projectile{X: x, Y: y, Xdirection: xd, Ydirection: yd, Image: image}
}

func (p *Projectile) Update(game *Game) error {
	p.X += p.Xdirection
	p.Y += p.Ydirection

	return nil
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Scale(0.5, 0.5)
	op.GeoM.Translate(p.X, p.Y)
	op.ColorM.RotateHue(10)
	screen.DrawImage(p.Image, op)
}
