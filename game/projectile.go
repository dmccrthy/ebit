package game

import (
	"github.com/dmccrthy/ebit/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type Projectile struct {
	X      float64
	Y      float64
	Vector *utils.Vector
	Width  int
	Height int
	Image  *ebiten.Image
}

func NewProjectile(startX, startY, endX, endY, speed float64, image *ebiten.Image) *Projectile {
	vector := &utils.Vector{X: endX - startX, Y: endY - startY}
	vector.Scale(speed)

	return &Projectile{
		X:      startX,
		Y:      startY,
		Vector: vector,
		Image:  image,
	}
}

func (p *Projectile) Update() {
	p.X += p.Vector.X
	p.Y += p.Vector.Y
}

func (p *Projectile) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Image, op)
}
