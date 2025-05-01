package game_manager

import (
	"GoPhysics/assets/ball"
	"GoPhysics/engine/colliders"
	"GoPhysics/engine/time_manager"
	"GoPhysics/engine/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

const fixedDeltaTime = 1.0 / 60.0

var Game = GameManager{}
var UpdatableObjects = []UpdatableObject{
	ball.NewBall(utils.Vector{X: 10, Y: 0}, true),
	ball.NewBall(utils.Vector{X: 200, Y: 0}, false),
}

type GameManager struct {
}

func (g *GameManager) Update() error {
	time_manager.Time.Update()

	for time_manager.Time.Accumulator >= fixedDeltaTime {
		g.FixedUpdate()
		time_manager.Time.Accumulator -= fixedDeltaTime
	}

	for _, updatableObject := range UpdatableObjects {
		updatableObject.Update()
	}

	return nil
}

func (g *GameManager) Draw(screen *ebiten.Image) {

	for _, updatableObject := range UpdatableObjects {
		updatableObject.Draw(screen)
	}

}

func (g *GameManager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *GameManager) FixedUpdate() {
	colliders.CM.CheckCollisions()

	for _, updatableObject := range UpdatableObjects {
		updatableObject.FixedUpdate()
	}
}
