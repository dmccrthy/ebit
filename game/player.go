package game

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	X      float64
	Y      float64
	Width  int
	Height int
	Image  *ebiten.Image
}

func (p *Player) Update(game *Game) error {

	// floor := float64(game.screenHeight - p.Height)

	// if p.Y < floor {
	// 	p.YVelocity += 0.1
	// }

	if ebiten.IsKeyPressed(ebiten.KeyW) || ebiten.IsKeyPressed(ebiten.KeyUp) {
		p.Y -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) || ebiten.IsKeyPressed(ebiten.KeyLeft) {
		p.X -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) || ebiten.IsKeyPressed(ebiten.KeyDown) {
		p.Y += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) || ebiten.IsKeyPressed(ebiten.KeyRight) {
		p.X += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		log.Print("Shutting Down")
		os.Exit(0)
	}

	return nil
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// op.GeoM.Scale(0.1, 0.1)
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Image, op)
}
