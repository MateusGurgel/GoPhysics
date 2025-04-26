package example

import (
	"GoPhysics/engine"
	"GoPhysics/engine/utils"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"log"
)

type Ball struct {
	object *engine.Object
}

func getBallImage() *ebiten.Image {
	ballImg, _, err := ebitenutil.NewImageFromFile("resources/Ball.png")

	if err != nil {
		log.Fatal(err)
	}

	return ballImg
}

func NewBall() *Ball {

	object := engine.NewObject(
		&utils.Vector{X: 120, Y: 120},
		nil,
		nil,
	)

	rb := engine.NewRigidBody(object, 1, 0.1)

	object.SetRigidBody(rb)

	return &Ball{
		object: object,
	}
}

func (b *Ball) FixedUpdate() {
	b.object.FixedUpdate()
	b.object.RigidBody.AddForce(utils.Vector{X: 0.1, Y: 0})
}

func (b *Ball) Update() {
	b.object.Update()
}

func (b *Ball) Draw(screen *ebiten.Image) {
	img := getBallImage()
	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Scale(1, 1)
	opts.GeoM.Translate(b.object.Position.X, b.object.Position.Y)
	screen.DrawImage(img, opts)
}
