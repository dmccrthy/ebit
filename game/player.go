package game

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	*Game
	X            float64
	Y            float64
	Width        int
	Height       int
	Image        *ebiten.Image
	ShotCooldown int
	ShotRate     int
}

func NewPlayer(game *Game) *Player {
	p := &Player{Game: game, ShotRate: 30}
	p.Width = 16
	p.Height = 16
	p.Image = p.Game.LoadImage("assets/player.png")

	return p
}

func (p *Player) Update() {
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

	// handle shooting
	if p.ShotCooldown > 0 {
		p.ShotCooldown--
	}

	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) && p.ShotCooldown == 0 {
		x := p.X + float64(p.Width-4)/2
		y := p.Y + float64(p.Height-4)/2
		img := p.Game.LoadImage("assets/bullet.png")
		mouseX, mouseY := ebiten.CursorPosition()

		p.Game.projectiles = append(p.Game.projectiles, NewProjectile(x, y, float64(mouseX), float64(mouseY), 2, img))
		p.ShotCooldown = p.ShotRate
	}

	if ebiten.IsKeyPressed(ebiten.KeyEscape) {
		log.Print("Shutting Down")
		os.Exit(0)
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	// op.GeoM.Scale(0.1, 0.1)
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Image, op)
}
