package main

import (
	"GoPhysics/engine/game_manager"
	"github.com/hajimehoshi/ebiten/v2"
	"log"
)

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&game_manager.Game); err != nil {
		log.Fatal(err)
	}
}
