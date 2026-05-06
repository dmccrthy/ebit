package game

import (
	"embed"
	"fmt"
	"image"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	fs           embed.FS
	screenWidth  int
	screenHeight int
	player       Player
	projectiles  []*Projectile
}

func NewGame(fs embed.FS) *Game {
	return &Game{fs: fs}
}

func (g *Game) Update() error {
	g.player.Update(g)

	for _, p := range g.projectiles {
		p.Update(g)
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{100, 100, 100, 0})
	
	g.player.Draw(screen)

	for _, p := range g.projectiles {
		p.Draw(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.screenWidth, g.screenHeight
}

func (g *Game) loadImage(path string) *ebiten.Image {
	file, err := g.fs.Open(path)
	if err != nil {
		log.Panic(err)
	}

	image, _, err := image.Decode(file)
	if err != nil {
		log.Panic(err)
	}

	return ebiten.NewImageFromImage(image)
}

func (g *Game) Init() {
	g.screenWidth, g.screenHeight = 128, 72
	g.player = Player{}
	g.player.Image = g.loadImage("assets/player.png")
	g.player.Width = 16
	g.player.Height = 16

	x, y := ebiten.Monitor().Size()
	fmt.Printf("%d\n", (x/y)*2)

	ebiten.SetWindowSize(ebiten.Monitor().Size())
	ebiten.SetFullscreen(true)
	ebiten.SetWindowTitle("Game")
}

// reset when player dies
func (g *Game) Reset() {

}
