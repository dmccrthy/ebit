package game

import (
	"log"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type Player struct {
	*Game
	X      float64
	Y      float64
	Width  int
	Height int
	Image  *ebiten.Image
}

func NewPlayer(game *Game) *Player {
	p := &Player{Game: game}
	p.Width = 16
	p.Height = 16
	p.Image = p.Game.LoadImage("assets/player.png")

	return p
}

func (p *Player) Update() error {
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

	x := p.X + float64(p.Width-4)/2
	y := p.Y + float64(p.Height-4)/2
	img := p.Game.LoadImage("assets/bullet.png")

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) {
		p.Game.projectiles = append(p.Game.projectiles, NewProjectile(x, y, 0, -2, img))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) {
		p.Game.projectiles = append(p.Game.projectiles, NewProjectile(x, y, -2, 0, img))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) {
		p.Game.projectiles = append(p.Game.projectiles, NewProjectile(x, y, 0, 2, img))
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) {
		p.Game.projectiles = append(p.Game.projectiles, NewProjectile(x, y, 2, 0, img))
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
