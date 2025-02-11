package scene02Game

import (
  graphics "github.com/quasilyte/ebitengine-graphics"
  "github.com/quasilyte/gmath"
  "github.com/quasilyte/gsignal"
)

type pickupNode struct {
  position gmath.Vec
  rect     *graphics.Rect
  scene    *scene
  score    int
  disposed bool

  EventDestroyed gsignal.Event[int]
}

func newPickupNode(position gmath.Vec) *pickupNode {
  return &pickupNode{position: position}
}

func (pickupNode *pickupNode) Init(scene *scene) {
  pickupNode.scene = scene
  context := scene.Controller().context
  pickupNode.score = context.Rand.IntRange(5, 10)
  pickupNode.rect = context.NewRect(16, 16)
  pickupNode.rect.Pos.Base = &pickupNode.position
  pickupNode.rect.SetFillColorScale(graphics.ColorScaleFromRGBA(200, 200, 0, 255))
  scene.AddGraphics(pickupNode.rect)
}

func (pickupNode *pickupNode) IsDisposed() bool {
  return pickupNode.disposed
}

func (pickupNode *pickupNode) Update(delta float64) {
  apple := pickupNode.scene.Controller().state.Apple
  if apple.position.DistanceTo(pickupNode.position) < 16 {
    pickupNode.pickUp()
  }
}

func (pickupNode *pickupNode) pickUp() {
  pickupNode.EventDestroyed.Emit(pickupNode.score)
  pickupNode.dispose()
}

func (pickupNode *pickupNode) dispose() {
  pickupNode.rect.Dispose()
  pickupNode.disposed = true
}
