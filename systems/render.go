package systems

import (
	"github.com/hpdobrica/netcode-experiments-client/components"

	"github.com/EngoEngine/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type renderSystemEntity struct {
	*ecs.BasicEntity
	*components.RenderComponent
	*components.LocationComponent
}

type RenderSystem struct {
	entities []renderSystemEntity
}

type Renderable interface {
	ecs.BasicFace
	components.RenderFace
	components.LocationFace
}

func (r *RenderSystem) AddByInterface(o ecs.Identifier) {
	obj := o.(Renderable)
	r.Add(obj.GetBasicEntity(), obj.GetRenderComponent(), obj.GetLocationComponent())
}

func (r *RenderSystem) GetEntities() []renderSystemEntity {
	return r.entities
}

func (r *RenderSystem) Add(basic *ecs.BasicEntity, rc *components.RenderComponent, lc *components.LocationComponent) {
	r.entities = append(r.entities, renderSystemEntity{basic, rc, lc})
}

func (r *RenderSystem) Remove(basic ecs.BasicEntity) {
	var delete int = -1
	for index, entity := range r.entities {
		if entity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		r.entities = append(r.entities[:delete], r.entities[delete+1:]...)
	}
}

func (r *RenderSystem) Update(_ float32) {}

func (r *RenderSystem) Render(screen *ebiten.Image) {
	for _, entity := range r.entities {

		opts := &ebiten.DrawImageOptions{}

		opts.GeoM.Translate(entity.GetLocationComponent().X, entity.GetLocationComponent().Y)
		// opts.GeoM.Scale(1.5, 1)
		// opts.GeoM.Rotate(0.1)

		screen.DrawImage(entity.RenderComponent.Img, opts)
		// fmt.Println(2, time.Now(), ebiten.IsKeyPressed(ebiten.KeyS))
	}
}
