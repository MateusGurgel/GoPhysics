package colliders

type CircleCollider struct {
	Radius    float64
	Observers []Observer
	Object    CollidableObject
}

func NewCircleCollider(object CollidableObject, radius float64, observer []Observer) *CircleCollider {

	collider := &CircleCollider{Radius: radius, Object: object, Observers: observer}

	CM.AddCollider(collider)

	return collider
}

func (c *CircleCollider) CheckCollisionWithCircle(Other *CircleCollider) bool {

	quadraticDistance := c.Object.GetPosition().GetEuclideanQuadraticDistance(Other.Object.GetPosition())
	radiusSum := c.Radius + Other.Radius

	return quadraticDistance < radiusSum*radiusSum
}

func (c *CircleCollider) TriggerCollision() {
	for _, observer := range c.Observers {
		observer.OnTrigger()
	}
}

func (c *CircleCollider) AddObserver(observer Observer) {
	c.Observers = append(c.Observers, observer)
}
