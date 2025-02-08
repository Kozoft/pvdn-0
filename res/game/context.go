package game

import (
  graphics "github.com/quasilyte/ebitengine-graphics"
  input "github.com/quasilyte/ebitengine-input"
  resource "github.com/quasilyte/ebitengine-resource"
  "github.com/quasilyte/gmath"
  "github.com/quasilyte/gscene"
)

type Context struct {
  Input         *input.Handler
  Loader        *resource.Loader
  Rand          gmath.Rand
  graphicsCache *graphics.Cache
  scene         gscene.GameRunner

  WindowWidth  int
  WindowHeight int
}

func ChangeScene[ControllerAccessor any](context *Context, controller gscene.Controller[ControllerAccessor]) {
  s := gscene.NewRootScene[ControllerAccessor](controller)
  context.scene = s
}

func NewContext() *Context {
  return &Context{
    graphicsCache: graphics.NewCache(),
  }
}

func (context *Context) NewRect(width, height float64) *graphics.Rect {
  return graphics.NewRect(context.graphicsCache, width, height)
}

func (context *Context) NewLabel(id resource.FontID) *graphics.Label {
  font := context.Loader.LoadFont(id)
  return graphics.NewLabel(context.graphicsCache, font.Face)
}

func (context *Context) NewSprite(id resource.ImageID) *graphics.Sprite {
  sprite := graphics.NewSprite(context.graphicsCache)
  if id == 0 {
    return sprite
  }
  image := context.Loader.LoadImage(id)
  sprite.SetImage(image.Data)
  if image.DefaultFrameWidth != 0 || image.DefaultFrameHeight != 0 {
    width, height := sprite.GetFrameSize()
    if image.DefaultFrameWidth != 0 {
      width = int(image.DefaultFrameWidth)
    }
    if image.DefaultFrameHeight != 0 {
      height = int(image.DefaultFrameHeight)
    }
    sprite.SetFrameSize(width, height)
  }
  return sprite
}

func (context *Context) CurrentScene() gscene.GameRunner {
  return context.scene
}
