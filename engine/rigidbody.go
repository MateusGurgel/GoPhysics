package engine

import (
	"GoPhysics/engine/time_manager"
	"GoPhysics/engine/utils"
)

type RigidBody struct {
	object       *Object
	velocity     *utils.Vector
	acceleration *utils.Vector
	mass         float64
	airDrag      float64
}

func NewRigidBody(object *Object, mass float64, airDrag float64) *RigidBody {
	return &RigidBody{object: object, mass: mass, airDrag: airDrag, velocity: &utils.Vector{}}
}

func (rb *RigidBody) AddForce(force utils.Vector) {
	accelerationProduct := force.Divide(rb.mass)
	deltaTime := time_manager.Time.DeltaTime

	velocityProduct := accelerationProduct.MultiplyByScalar(deltaTime)

	newPosition := rb.object.Position.Add(velocityProduct)
	rb.object.SetPosition(newPosition)
}

func (rb *RigidBody) ApplyVelocityOnObject() {
	newPosition := rb.object.Position.Add(rb.velocity)
	rb.object.SetPosition(newPosition)
}

func (rb *RigidBody) UpdatePhysics() {
	rb.UpdatePhysics()
}
