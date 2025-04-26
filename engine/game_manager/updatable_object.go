package game_manager

import "github.com/hajimehoshi/ebiten/v2"

type UpdatableObject interface {
	Update()
	Draw(screen *ebiten.Image)
	FixedUpdate()
}
