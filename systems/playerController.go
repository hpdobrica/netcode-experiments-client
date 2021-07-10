package systems

import (
	"github.com/hpdobrica/netcode-experiments-client/components"

	"github.com/EngoEngine/ecs"
	"github.com/hajimehoshi/ebiten/v2"
)

type playerControllerSystemEntity struct {
	*ecs.BasicEntity
	*components.PlayerControllerComponent
	*components.LocationComponent
}

type PlayerControllerSystem struct {
	entities []playerControllerSystemEntity
}

type PlayerControllable interface {
	ecs.BasicFace
	components.PlayerControllerFace
	components.LocationFace
}

func (p *PlayerControllerSystem) AddByInterface(o ecs.Identifier) {
	obj := o.(PlayerControllable)
	p.Add(obj.GetBasicEntity(), obj.GetPlayerControllerComponent(), obj.GetLocationComponent())
}

func (p *PlayerControllerSystem) GetEntities() []playerControllerSystemEntity {
	return p.entities
}

func (p *PlayerControllerSystem) Add(basic *ecs.BasicEntity, pc *components.PlayerControllerComponent, lc *components.LocationComponent) {
	p.entities = append(p.entities, playerControllerSystemEntity{basic, pc, lc})
}

func (p *PlayerControllerSystem) Remove(basic ecs.BasicEntity) {
	var delete int = -1
	for index, entity := range p.entities {
		if entity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		p.entities = append(p.entities[:delete], p.entities[delete+1:]...)
	}
}

func (p *PlayerControllerSystem) Update(dt float32) {
	speed := float64(5)
	for _, entity := range p.entities {
		loc := entity.GetLocationComponent()
		if ebiten.IsKeyPressed(ebiten.KeyW) {
			loc.Y -= speed
		}

		if ebiten.IsKeyPressed(ebiten.KeyS) {
			loc.Y += speed
		}

		if ebiten.IsKeyPressed(ebiten.KeyA) {
			loc.X -= speed
		}

		if ebiten.IsKeyPressed(ebiten.KeyD) {
			loc.X += speed
		}

	}
}
