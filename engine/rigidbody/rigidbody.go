package rigidbody

import (
	"GoPhysics/engine/time_manager"
	"GoPhysics/engine/utils"
)

type RigidBody struct {
	object   PhysicsObject
	velocity *utils.Vector
	mass     float64
	airDrag  float64
}

func NewRigidBody(object PhysicsObject, mass float64, airDrag float64) *RigidBody {
	return &RigidBody{object: object, mass: mass, airDrag: airDrag, velocity: &utils.Vector{}}
}

func (rb *RigidBody) SetVelocity(velocity *utils.Vector) {
	rb.velocity = velocity
}

func (rb *RigidBody) AddForce(force utils.Vector) {
	accelerationProduct := force.Divide(rb.mass)
	deltaTime := time_manager.Time.DeltaTime

	velocityProduct := accelerationProduct.MultiplyByScalar(deltaTime)
	velocityProduct = velocityProduct.MultiplyByScalar(0.03)
	finalVelocity := velocityProduct.Add(rb.velocity)

	rb.SetVelocity(finalVelocity)
}

func (rb *RigidBody) AddGravity() {
	gravityForce := utils.Vector{Y: rb.mass * 9.8}
	rb.AddForce(gravityForce)
}

func (rb *RigidBody) ApplyVelocityOnObject() {
	newPosition := rb.object.GetPosition().Add(rb.velocity)
	rb.object.SetPosition(newPosition)
}

func (rb *RigidBody) UpdatePhysics() {
	rb.AddGravity()
	rb.ApplyVelocityOnObject()
}
