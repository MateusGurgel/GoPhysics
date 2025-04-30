package colliders

type ColliderManager struct {
	Colliders []*CircleCollider
	Object    CollidableObject
}

func NewColliderManager() *ColliderManager {
	return &ColliderManager{}
}

func (cm *ColliderManager) AddCollider(collider *CircleCollider) {
	cm.Colliders = append(cm.Colliders, collider)
}

func (cm *ColliderManager) CheckCollisions() {
	for i, collider := range cm.Colliders {
		for j, other := range cm.Colliders {

			if i == j {
				continue
			}

			if collider.CheckCollisionWithCircle(other) {
				collider.TriggerCollision()
			}
		}
	}
}

var CM = NewColliderManager()
