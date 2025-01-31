package main

import (
  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/audio"
  input "github.com/quasilyte/ebitengine-input"
  resource "github.com/quasilyte/ebitengine-resource"
  "github.com/quasilyte/gmath"
  "pvdn-0/res/assets"
  "pvdn-0/res/controls"
  "pvdn-0/res/game"
)

type Player struct {
  pos gmath.Vec // {X, Y}
  img *ebiten.Image
}

type aGame struct {
  inputSystem input.System
  context     *game.Context
  player      *Player
}

func main() {
  context := &game.Context{
    Loader:       createLoader(),
    WindowWidth:  320,
    WindowHeight: 240,
  }
  theGame := &aGame{
    context: context,
  }
  theGame.inputSystem.Init(input.SystemConfig{
    DevicesEnabled: input.AnyDevice,
  })
  context.Input = theGame.inputSystem.NewHandler(0, controls.DefaultKeymap)
  ebiten.SetWindowSize(theGame.context.WindowWidth, theGame.context.WindowHeight)
  ebiten.SetWindowTitle("Ebitengine Apple Sample")
  assets.RegisterResources(context.Loader)
  theGame.init()
  if err := ebiten.RunGame(theGame); err != nil {
    panic(err)
  }
}

// Update game logic here
func (localGame *aGame) Update() error {
  localGame.inputSystem.Update()
  speed := 64.0 * (1.0 / 60)
  var v gmath.Vec
  if localGame.context.Input.ActionIsPressed(controls.ActionMoveRight) {
    v.X += speed
  }
  if localGame.context.Input.ActionIsPressed(controls.ActionMoveDown) {
    v.Y += speed
  }
  if localGame.context.Input.ActionIsPressed(controls.ActionMoveLeft) {
    v.X -= speed
  }
  if localGame.context.Input.ActionIsPressed(controls.ActionMoveUp) {
    v.Y -= speed
  }
  localGame.player.pos = localGame.player.pos.Add(v)
  return nil
}

// Draw drawing here
func (localGame *aGame) Draw(screen *ebiten.Image) {
  var options ebiten.DrawImageOptions
  options.GeoM.Translate(localGame.player.pos.X, localGame.player.pos.Y)
  screen.DrawImage(localGame.player.img, &options)
}

func (localGame *aGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
  return localGame.context.WindowWidth, localGame.context.WindowHeight
}

func (localGame *aGame) init() {
  apple := localGame.context.Loader.LoadImage(assets.ImageApple).Data
  localGame.player = &Player{img: apple}
}

func createLoader() *resource.Loader {
  sampleRate := 44100
  audioContext := audio.NewContext(sampleRate)
  loader := resource.NewLoader(audioContext)
  loader.OpenAssetFunc = assets.OpenAsset
  return loader
}
