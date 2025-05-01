package ball

import (
	"GoPhysics/engine"
	"GoPhysics/engine/colliders"
	"GoPhysics/engine/rigidbody"
	"GoPhysics/engine/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Ball struct {
	Position  *utils.Vector
	RigidBody *rigidbody.RigidBody
	Colliders []engine.Collider
	forceLeft bool
}

func NewBall(position utils.Vector, forceLeft bool) *Ball {

	object := &Ball{
		&position,
		nil,
		nil,
		forceLeft,
	}

	rb := rigidbody.NewRigidBody(object, 20, 0.1, false)
	collider := colliders.NewCircleCollider(object, 10)

	object.RigidBody = rb
	object.Colliders = append(object.Colliders, collider)

	return object
}

func (o *Ball) Update() {
}

func (o *Ball) FixedUpdate() {
	o.RigidBody.UpdatePhysics()

	if o.forceLeft {
		o.RigidBody.AddForce(utils.Vector{X: 15, Y: 0})
	} else {
		o.RigidBody.AddForce(utils.Vector{X: -15, Y: 0})
	}

}

func (o *Ball) Draw(screen *ebiten.Image) {
	img := getBallImage()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(1, 1)
	opts.GeoM.Translate(o.Position.X, o.Position.Y)
	screen.DrawImage(img, opts)
}

func getBallImage() *ebiten.Image {
	ballImg, _, err := ebitenutil.NewImageFromFile("resources/Ball.png")

	if err != nil {
		log.Fatal(err)
	}

	return ballImg
}

func (o *Ball) GetPosition() *utils.Vector {
	return o.Position
}

func (o *Ball) SetPosition(position *utils.Vector) {
	o.Position = position
}
