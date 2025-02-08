package scenes

import (
  "github.com/quasilyte/gmath"
  "github.com/quasilyte/gscene"
  "pvdn-0/res/game"
)

type scene = gscene.Scene[*WalksceneController]

type WalksceneController struct {
  context *game.Context
  scene   *gscene.RootScene[*WalksceneController]
}

func NewWalksceneController(context *game.Context) *WalksceneController {
  return &WalksceneController{context: context}
}

func (controller *WalksceneController) Init(rootScene *gscene.RootScene[*WalksceneController]) {
  controller.scene = rootScene
  apple := newAppleNode(gmath.Vec{X: 64, Y: 64})
  rootScene.AddObject(apple)
}

func (controller *WalksceneController) Update(delta float64) {
}
