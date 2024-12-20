package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	resource "github.com/quasilyte/ebitengine-resource"
	"github.com/quasilyte/gmath"
	"pvdn-0/res/assets"
)

type Player struct {
	pos gmath.Vec // {X, Y}
	img *ebiten.Image
}

type myGame struct {
	windowWidth  int
	windowHeight int
	loader       *resource.Loader
	player       *Player
}

func main() {
	g := &myGame{
		windowWidth:  320,
		windowHeight: 240,
		loader:       createLoader(),
	}
	ebiten.SetWindowSize(g.windowWidth, g.windowHeight)
	ebiten.SetWindowTitle("Ebitengine Apple Sample")
	assets.RegisterResources(g.loader)
	g.init()
	if err := ebiten.RunGame(g); err != nil {
		panic(err)
	}
}

// Update game logic here
func (g *myGame) Update() error {
	g.player.pos.X += 16 * (1.0 / 60.0)
	return nil
}

// Draw drawing here
func (g *myGame) Draw(screen *ebiten.Image) {
	var options ebiten.DrawImageOptions
	options.GeoM.Translate(g.player.pos.X, g.player.pos.Y)
	screen.DrawImage(g.player.img, &options)
}

func (g *myGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.windowWidth, g.windowHeight
}

func (g *myGame) init() {
	apple := g.loader.LoadImage(assets.ImageApple).Data
	g.player = &Player{img: apple}
}

func createLoader() *resource.Loader {
	sampleRate := 44100
	audioContext := audio.NewContext(sampleRate)
	loader := resource.NewLoader(audioContext)
	loader.OpenAssetFunc = assets.OpenAsset
	return loader
}
