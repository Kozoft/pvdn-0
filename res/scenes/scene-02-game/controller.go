package scene02Game

import (
	"fmt"

	graphics "github.com/quasilyte/ebitengine-graphics"
	"github.com/quasilyte/gmath"
	"github.com/quasilyte/gscene"
	"pvdn-0/res/assets"
	"pvdn-0/res/controls"
	"pvdn-0/res/game"
)

type scene = gscene.Scene[*Controller02]
type Controller02 struct {
	context    *game.Context
	scene      *gscene.RootScene[*Controller02]
	state      *SceneState
	scoreLabel *graphics.Label
	score      int
}

func NewController02(context *game.Context) *Controller02 {
	return &Controller02{context: context}
}

func (controller *Controller02) Init(rootScene *gscene.RootScene[*Controller02]) {
	controller.scene = rootScene
	apple := newAppleNode(gmath.Vec{X: 64, Y: 64})
	rootScene.AddObject(apple)
	controller.state = &SceneState{Apple: apple}
	controller.scoreLabel = controller.context.CreateLabel(assets.FontNormal)
	controller.scoreLabel.Pos.Offset = gmath.Vec{X: 4, Y: 4}
	rootScene.AddGraphics(controller.scoreLabel)
	controller.createPickup()
	controller.addScore(0)
}

func (controller *Controller02) createPickup() {
	pickupNode := newPickupNode(gmath.Vec{
		X: controller.context.Rand.FloatRange(0, float64(controller.context.WindowWidth)),
		Y: controller.context.Rand.FloatRange(0, float64(controller.context.WindowHeight)),
	})
	pickupNode.EventDestroyed.Connect(nil, func(score int) {
		controller.addScore(score)
		controller.createPickup()
	})
	controller.scene.AddObject(pickupNode)
}

func (controller *Controller02) addScore(score int) {
	controller.score += score
	controller.scoreLabel.SetText(fmt.Sprintf("score: %d", controller.score))
}

func (controller *Controller02) Update(delta float64) {
	if controller.context.Input.ActionIsJustPressed(controls.ActionRestart) {
		game.ChangeScene(controller.context, NewController02(controller.context))
	}
}
