package engine

import (
	"GoPhysics/engine/utils"
)

type Object struct {
	Position  *utils.Vector
	RigidBody *RigidBody
	Colliders []Collider
}

func NewObject(position *utils.Vector, collider []Collider, rigidBody *RigidBody) *Object {
	return &Object{Position: position, Colliders: collider, RigidBody: rigidBody}
}

func (o *Object) SetRigidBody(rigidBody *RigidBody) {
	o.RigidBody = rigidBody
}

func (o *Object) SetPosition(position *utils.Vector) {
	o.Position = position
}

func (o *Object) AddCollider(collider Collider) {
	o.Colliders = append(o.Colliders, collider)
}

func (o *Object) GetPosition() *utils.Vector {
	return o.Position
}

func (o *Object) Update() {

}

func (o *Object) FixedUpdate() {

	if o.RigidBody != nil {
		o.RigidBody.UpdatePhysics()
	}
}
