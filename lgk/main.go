package main

import (
	"time"

	"github.com/codecat/go-enet"
	"github.com/codecat/go-libs/log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	input "github.com/quasilyte/ebitengine-input"
	resource "github.com/quasilyte/ebitengine-resource"
	"pvdn-0/res/assets"
	"pvdn-0/res/controls"
	"pvdn-0/res/game"
	scene01Splash "pvdn-0/res/scenes/scene-01-splash"
)

type aGame struct {
	inputSystem input.System
	client      enet.Host
	peer        enet.Peer
	context     *game.Context
}

func main() {
	context := game.CreateContext()
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

	enet.Initialize()

	theGame.client, err = enet.NewHost(nil, 1, 1, 0, 0)
	if err != nil {
		log.Error("Couldn't create host: %s", err.Error())
	}
	theGame.peer, err = theGame.client.Connect(enet.NewAddress("127.0.0.1", 8095), 1, 0)
	if err != nil {
		log.Error("Couldn't connect: %s", err.Error())
	}

	context.Input = theGame.inputSystem.NewHandler(0, controls.DefaultKeymap)
	ebiten.SetWindowSize(theGame.context.WindowWidth, theGame.context.WindowHeight)
	ebiten.SetWindowTitle("Ebitengine Apple Sample")
	assets.RegisterResources(context.Loader)
	game.ChangeScene(context, scene01Splash.NewController01(context))
	if err = ebiten.RunGame(theGame); err != nil {
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
