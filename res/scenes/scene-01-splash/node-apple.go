package scene01Splash

import (
  graphics "github.com/quasilyte/ebitengine-graphics"
  "github.com/quasilyte/gmath"
  "pvdn-0/res/assets"
)

type appleNode struct {
  position gmath.Vec
  sprite   *graphics.Sprite
}

func newAppleNode(position gmath.Vec) *appleNode {
  return &appleNode{
    position: position,
  }
}

func (apple *appleNode) Init(scene *scene) {
  context := scene.Controller().context
  apple.sprite = context.NewSprite(assets.ImageApple)
  apple.sprite.Pos.Base = &apple.position
  scene.AddGraphics(apple.sprite)
}

func (apple *appleNode) IsDisposed() bool {
  return false
}

func (apple *appleNode) Update(delta float64) {}
