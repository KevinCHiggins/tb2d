package tb2d

import (
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/img"
)


type Graphic interface {
	// this is just the dimensions, not viewport coords
	GetWidth() int
	GetHeight() int
	Blit(dest Graphic, clipRect Rect) // error?
	//GetSurf() *sdl.Surface // ruins the whole point, just for testing
}

// Exists to satisfy the Graphic interface (more precisely, a *sdlSurfaceWrapper does) while referencing an SDL Surface
type sdlSurfaceWrapper struct {
	surf *sdl.Surface
}

func (wrappedSurf *sdlSurfaceWrapper) GetWidth() int {
	return int(wrappedSurf.surf.W)
}
/*
func (graphic *sdlSurfaceWrapper) GetSurf() *sdl.Surface {
	return graphic.surf
}
*/

func (wrappedSurf *sdlSurfaceWrapper) GetHeight() int {
	return int(wrappedSurf.surf.H)
}

// should of course return an error
func NewGraphicFromFile(s string) Graphic {
	surf, err := img.Load(s)
	if err != nil {
		panic(err)
	}
	sg := sdlSurfaceWrapper{surf}
	g := Graphic(&sg)
	return g
}

func NewBlankGraphic(width, height int) Graphic {
	surf, err := sdl.CreateRGBSurfaceWithFormat(0, int32(width), int32(height), 32, pxfmt.Format)
	if err != nil {
		panic(err)
	}
	sg := sdlSurfaceWrapper{surf}
	return &sg

}

func (wrappedSurf *sdlSurfaceWrapper) Blit (dest Graphic, clipRect Rect) {
	destWrappedSurf := dest.(*sdlSurfaceWrapper)
	sdlRect := sdl.Rect{int32(clipRect.X), int32(clipRect.Y), int32(clipRect.W), int32(clipRect.H)}
	wrappedSurf.surf.Blit(&sdl.Rect{0, 0, int32(wrappedSurf.GetWidth()), int32(wrappedSurf.GetHeight())}, destWrappedSurf.surf, &sdlRect)

}
