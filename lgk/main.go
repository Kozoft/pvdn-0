package main

import (
  "time"

  "github.com/hajimehoshi/ebiten/v2"
  "github.com/hajimehoshi/ebiten/v2/audio"
  input "github.com/quasilyte/ebitengine-input"
  resource "github.com/quasilyte/ebitengine-resource"
  "pvdn-0/res/assets"
  "pvdn-0/res/controls"
  "pvdn-0/res/game"
  splash "pvdn-0/res/scenes/scenes-01-splash"
)

type aGame struct {
  inputSystem input.System
  context     *game.Context
}

func main() {
  context := game.NewContext()
  context.Loader = createLoader()
  context.WindowWidth = 640
  context.WindowHeight = 480
  context.Rand.SetSeed(time.Now().Unix())
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
  game.ChangeScene(context, splash.NewController01(context))
  if err := ebiten.RunGame(theGame); err != nil {
    panic(err)
  }
}

// Update game logic
func (localGame *aGame) Update() error {
  localGame.inputSystem.Update()
  localGame.context.CurrentScene().UpdateWithDelta(1.0 / 60.0)
  return nil
}

// Draw drawing
func (localGame *aGame) Draw(screen *ebiten.Image) {
  localGame.context.CurrentScene().Draw(screen)
}

func (localGame *aGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
  return localGame.context.WindowWidth, localGame.context.WindowHeight
}

func createLoader() *resource.Loader {
  sampleRate := 44100
  audioContext := audio.NewContext(sampleRate)
  loader := resource.NewLoader(audioContext)
  loader.OpenAssetFunc = assets.OpenAsset
  return loader
}
