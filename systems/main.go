package systems

import "github.com/hajimehoshi/ebiten/v2"

type Renderer interface {
	Render(screen *ebiten.Image)
}
