package engine

import (
	"GoPhysics/engine/utils"
	"github.com/hajimehoshi/ebiten/v2"
)

type ObjectInterface interface {
	Update()
	FixedUpdate()
	Draw(screen *ebiten.Image)
	GetPosition() *utils.Vector
	GetRigidBody() *RigidBody
	GetColliders() []Collider
}
