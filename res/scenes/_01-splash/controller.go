package scenes

import (
  graphics "github.com/quasilyte/ebitengine-graphics"
  "github.com/quasilyte/gscene"
  "pvdn-0/res/assets"
  "pvdn-0/res/controls"
  "pvdn-0/res/game"
  "pvdn-0/res/scenes/_02-game"
)

type SplashController struct {
  context *game.Context
}

func NewSplashController(context *game.Context) *SplashController {
  return &SplashController{context: context}
}

func (controller *SplashController) Init(scene *gscene.SimpleRootScene) {
  label := controller.context.NewLabel(assets.FontBig)
  label.SetAlignHorizontal(graphics.AlignHorizontalCenter)
  label.SetAlignVertical(graphics.AlignVerticalCenter)
  label.SetSize(controller.context.WindowWidth, controller.context.WindowHeight)
  label.SetText("Press [Enter] to continue")
  scene.AddGraphics(label)
}

func (controller *SplashController) Update(delta float64) {
  if controller.context.Input.ActionIsJustPressed(controls.ActionConfirm) {
    game.ChangeScene(controller.context, scenes.NewWalksceneController(controller.context))
  }
}
