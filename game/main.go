package game

import (
	"fmt"
	"image/color"
	_ "image/png"
	"log"

	"github.com/hpdobrica/netcode-experiments-client/components"
	"github.com/hpdobrica/netcode-experiments-client/entities"
	"github.com/hpdobrica/netcode-experiments-client/systems"

	"github.com/EngoEngine/ecs"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type Game struct {
	systems   []ecs.System
	renderers []systems.Renderer
	world     ecs.World
}

func NewGame() *Game {
	g := &Game{}
	g.Start()

	return g
}

func (g *Game) Start() {
	g.world = ecs.World{}

	var locationable *systems.Locationable
	lsys := systems.LocationSystem{}
	g.world.AddSystemInterface(&lsys, locationable, nil)

	var renderable *systems.Renderable
	rsys := systems.RenderSystem{}
	g.world.AddSystemInterface(&rsys, renderable, nil)

	var playerControllable *systems.PlayerControllable
	psys := systems.PlayerControllerSystem{}
	g.world.AddSystemInterface(&psys, playerControllable, nil)

	g.initEntities()

	for _, system := range g.world.Systems() {
		r, ok := system.(systems.Renderer)
		fmt.Println(r, ok)

		if ok {
			g.renderers = append(g.renderers, r)
		} else {
			g.systems = append(g.systems, system)
		}
	}

}

func (g *Game) Update() error {

	dt := 1.0 / float32(ebiten.MaxTPS())
	g.world.Update(dt)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{0xff, 0xff, 0, 0xff})
	ebitenutil.DebugPrint(screen, "Yello, World!")

	for _, s := range g.renderers {
		s.Render(screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	// return 320, 240
	return 640, 480
}

func (g *Game) initEntities() {
	img, _, err := ebitenutil.NewImageFromFile("assets/img/gopher.png")

	if err != nil {
		log.Fatal(err)
	}

	gopher := entities.Gopher{
		BasicEntity:               ecs.NewBasic(),
		LocationComponent:         &components.LocationComponent{X: 320, Y: 365},
		RenderComponent:           &components.RenderComponent{Img: img},
		PlayerControllerComponent: &components.PlayerControllerComponent{},
	}

	g.world.AddEntity(&gopher)

}
