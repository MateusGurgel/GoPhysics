package game_manager

import (
	"GoPhysics/engine/time_manager"
	"GoPhysics/example"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

const fixedDeltaTime = 1.0 / 60.0

var Game = GameManager{}
var Ball = example.NewBall()

type GameManager struct {
}

func (g *GameManager) Update() error {
	time_manager.Time.Update()

	for time_manager.Time.Accumulator >= fixedDeltaTime {
		g.FixedUpdate()
		time_manager.Time.Accumulator -= fixedDeltaTime
	}

	return nil
}

func (g *GameManager) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")
	Ball.Draw(screen)
}

func (g *GameManager) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func (g *GameManager) FixedUpdate() {
}
