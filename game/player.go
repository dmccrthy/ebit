package game

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	X      float64
	Y      float64
	Width  int
	Height int
	Image  *ebiten.Image
}

func (p *Player) Update(game *Game) error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		p.Y -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyA) {
		p.X -= 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyS) {
		p.Y += 1
	}

	if ebiten.IsKeyPressed(ebiten.KeyD) {
		p.X += 1
	}

	x := p.X + float64(p.Width - 8) / 2
	y := p.Y + float64(p.Height - 8) / 2

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		game.projectiles = append(game.projectiles, NewProjectile(x, y, 0, -1, game.loadImage("assets/player.png")))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		game.projectiles = append(game.projectiles, NewProjectile(x, y, -1, 0, game.loadImage("assets/player.png")))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		game.projectiles = append(game.projectiles, NewProjectile(x, y, 0, 1, game.loadImage("assets/player.png")))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		game.projectiles = append(game.projectiles, NewProjectile(x, y, 1, 0, game.loadImage("assets/player.png")))
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
