package engine

import (
	"GoPhysics/engine/utils"
)

type BaseCollider interface {
	CheckCollisionWithCircle(SelfPosition utils.Vector, TargetPosition utils.Vector, r float32) bool
}
