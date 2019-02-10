package s6d

import (
	"fmt"
	"image"

	"golang.org/x/exp/shiny/screen"
)

// term implements screen.Screen
type term struct{}

func (s term) NewBuffer(size image.Point) (screen.Buffer, error) {
	m := image.NewRGBA(image.Rectangle{Max: size})
	return buffer(*m), nil
}
func (s term) NewTexture(size image.Point) (screen.Texture, error) {
	return nil, fmt.Errorf("s6d: textures are not supported")
}
func (s term) NewWindow(opts *screen.NewWindowOptions) (screen.Window, error) {
	return window{}, nil
}

type buffer image.RGBA

func (b buffer) Release()                {}
func (b buffer) Size() image.Point       { return b.Rect().Max }
func (b buffer) RGBA() *image.RGBA       { return *b }
func (b buffer) Bounds() image.Rectangle { return image.RGBA(b).Bounds() }
