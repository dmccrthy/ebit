package main

import (
	"embed"
	"log"

	"github.com/dmccrthy/ebit/game"
	"github.com/hajimehoshi/ebiten/v2"
)

//go:embed assets
var fs embed.FS

func main() {
	g := game.NewGame(fs)
	g.Init()

	err := ebiten.RunGame(g)
	if err != nil {
		log.Fatal(err)
	}
}
