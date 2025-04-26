package colliders

import "GoPhysics/engine/utils"

type CircleCollider struct {
	Radius float64
}

func (c *CircleCollider) CheckCollisionWithCircle(SelfPosition utils.Vector, OtherPosition *utils.Vector, Other *CircleCollider) bool {
	quadraticDistance := SelfPosition.GetEuclideanQuadraticDistance(OtherPosition)
	radiusSum := c.Radius + Other.Radius

	return quadraticDistance < radiusSum*radiusSum
}
