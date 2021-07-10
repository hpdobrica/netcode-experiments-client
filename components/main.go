package components

import (
	_ "image/png"

	"github.com/hajimehoshi/ebiten/v2"
)

type LocationComponent struct {
	X float64
	Y float64
}

func (l *LocationComponent) GetLocationComponent() *LocationComponent {
	return l
}

type LocationFace interface {
	GetLocationComponent() *LocationComponent
}

type RenderComponent struct {
	Img    *ebiten.Image
	Hidden bool
}

func (l *RenderComponent) GetRenderComponent() *RenderComponent {
	return l
}

type RenderFace interface {
	GetRenderComponent() *RenderComponent
}

type PlayerControllerComponent struct {
}

func (p *PlayerControllerComponent) GetPlayerControllerComponent() *PlayerControllerComponent {
	return p
}

type PlayerControllerFace interface {
	GetPlayerControllerComponent() *PlayerControllerComponent
}
