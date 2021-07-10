package main

import (
	_ "image/png"
	"log"

	"github.com/hpdobrica/netcode-experiments-client/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	g := game.NewGame()
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
