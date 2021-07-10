package entities

import (
	"github.com/hpdobrica/netcode-experiments-client/components"

	"github.com/EngoEngine/ecs"
)

type Gopher struct {
	ecs.BasicEntity
	*components.LocationComponent
	*components.RenderComponent
	*components.PlayerControllerComponent
}
