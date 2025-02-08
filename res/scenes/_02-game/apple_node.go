package scenes

import (
  graphics "github.com/quasilyte/ebitengine-graphics"
  input "github.com/quasilyte/ebitengine-input"
  "github.com/quasilyte/gmath"
  "pvdn-0/res/assets"
  "pvdn-0/res/controls"
)

type appleNode struct {
  input    *input.Handler
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

  apple.input = context.Input

  apple.sprite = context.NewSprite(assets.ImageApple)
  apple.sprite.Pos.Base = &apple.position
  scene.AddGraphics(apple.sprite)
}

func (apple *appleNode) IsDisposed() bool {
  return false
}

func (apple *appleNode) Update(delta float64) {
  speed := 64.0 * (1.0 / 60)
  var vector gmath.Vec
  if apple.input.ActionIsPressed(controls.ActionMoveRight) {
    vector.X += speed
  }
  if apple.input.ActionIsPressed(controls.ActionMoveDown) {
    vector.Y += speed
  }
  if apple.input.ActionIsPressed(controls.ActionMoveLeft) {
    vector.X -= speed
  }
  if apple.input.ActionIsPressed(controls.ActionMoveUp) {
    vector.Y -= speed
  }
  apple.position = apple.position.Add(vector)
}
