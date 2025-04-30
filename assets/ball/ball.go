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
}

func NewBall() *Ball {

	object := &Ball{
		&utils.Vector{X: 120, Y: 0},
		nil,
		nil,
	}

	rb := rigidbody.NewRigidBody(object, 20, 0.1, false)
	collider := colliders.NewCircleCollider(10)

	object.RigidBody = rb
	object.Colliders = append(object.Colliders, collider)

	return object
}

func (o *Ball) Update() {
}

func (o *Ball) FixedUpdate() {
	o.RigidBody.UpdatePhysics()
	o.RigidBody.AddForce(utils.Vector{X: 0.1, Y: 0})
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
