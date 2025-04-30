package engine

import (
	"GoPhysics/engine/colliders"
)

type Collider interface {
	CheckCollisionWithCircle(Other *colliders.CircleCollider) bool
}
