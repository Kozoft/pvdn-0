package scene01Splash

import (
  graphics "github.com/quasilyte/ebitengine-graphics"
  "github.com/quasilyte/gmath"
  "github.com/quasilyte/gscene"
  "pvdn-0/res/assets"
  "pvdn-0/res/controls"
  "pvdn-0/res/game"
  scene02Game "pvdn-0/res/scenes/scene-02-game"
)

type scene = gscene.Scene[*Controller01]
type Controller01 struct {
  context *game.Context
  scene   *gscene.RootScene[*Controller01]
}

func NewController01(context *game.Context) *Controller01 {
  return &Controller01{context: context}
}

func (controller *Controller01) Init(scene *gscene.RootScene[*Controller01]) {
  controller.scene = scene
  apple := newAppleNode(gmath.Vec{X: 128, Y: 128})
  scene.AddObject(apple)
  label := controller.context.NewLabel(assets.FontBig)
  label.SetAlignHorizontal(graphics.AlignHorizontalCenter)
  label.SetAlignVertical(graphics.AlignVerticalCenter)
  label.SetSize(controller.context.WindowWidth, controller.context.WindowHeight)
  label.SetText("Bitte Knopf drucken")
  scene.AddGraphics(label)
}

func (controller *Controller01) Update(delta float64) {
  if controller.context.Input.ActionIsJustPressed(controls.ActionConfirm) {
    game.ChangeScene(controller.context, scene02Game.NewController02(controller.context))
  }
}
