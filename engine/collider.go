package engine

import (
	"GoPhysics/engine/colliders"
	"GoPhysics/engine/utils"
)

type Collider interface {
	CheckCollisionWithCircle(SelfPosition *utils.Vector, OtherPosition *utils.Vector, Other *colliders.CircleCollider) bool
}
