package systems

import (
	"github.com/hpdobrica/netcode-experiments-client/components"

	"github.com/EngoEngine/ecs"
)

type locationSystemEntity struct {
	*ecs.BasicEntity
	*components.LocationComponent
}

type LocationSystem struct {
	entities []locationSystemEntity
}

type Locationable interface {
	ecs.BasicFace
	components.LocationFace
}

func (l *LocationSystem) AddByInterface(o ecs.Identifier) {
	obj := o.(Locationable)
	l.Add(obj.GetBasicEntity(), obj.GetLocationComponent())
}

func (l *LocationSystem) GetEntities() []locationSystemEntity {
	return l.entities
}

func (l *LocationSystem) Add(basic *ecs.BasicEntity, lc *components.LocationComponent) {
	l.entities = append(l.entities, locationSystemEntity{basic, lc})
}

func (l *LocationSystem) Remove(basic ecs.BasicEntity) {
	var delete int = -1
	for index, entity := range l.entities {
		if entity.ID() == basic.ID() {
			delete = index
			break
		}
	}
	if delete >= 0 {
		l.entities = append(l.entities[:delete], l.entities[delete+1:]...)
	}
}

func (l *LocationSystem) Update(dt float32) {
	// for _, entity := range l.entities {
	// 	entity.GetLocationComponent()
	// 	// fmt.Println("I would like to tell you", entity.ID(), "that it has been", dt, "seconds since the last time we spoke. Your X is", entity.LocationComponent.X)
	// }
}
