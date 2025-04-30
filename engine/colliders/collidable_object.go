package colliders

import (
	"GoPhysics/engine/utils"
)

type CollidableObject interface {
	GetPosition() *utils.Vector
	SetPosition(position *utils.Vector)
}
