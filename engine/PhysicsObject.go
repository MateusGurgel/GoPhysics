package engine

import (
	"GoPhysics/engine/utils"
)

type PhysicsObject interface {
	GetPosition() *utils.Vector
	SetPosition(position *utils.Vector)
}
