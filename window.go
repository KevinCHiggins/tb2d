package tb2d
import (
	"github.com/veandco/go-sdl2/sdl"
)

var viewport Graphic
var window *sdl.Window
var pxfmt *sdl.PixelFormat

//func SetUpWindow(w, h int, fullscreen bool, g Graphic) {
func SetUpWindow(w, h int, fullscreen bool) {
	var err error
	if err = sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}

	// assign to global
	window, err = sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		int32(w), int32(h), sdl.WINDOW_SHOWN)
	if err != nil {
		panic(err)
	}
	


	if fullscreen {
		window.SetFullscreen(sdl.WINDOW_FULLSCREEN)
	}
	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	// assign to global so all gfx routines use same format
	pxfmt = surface.Format

	surface.FillRect(nil, 0)

	rect := sdl.Rect{0, 0, 200, 200}
	surface.FillRect(&rect, 0xffff0000)


	viewport = &sdlSurfaceWrapper{surface}

	/*

	s := g.GetSurf()
	rect = sdl.Rect{0, 0, 120, 36}
	//s.FillRect(&rect, 0xffffff00)
	print(s.W)
	print(s.H)
	print(s.Format.Format == surface.Format.Format)
	print(s.Format.Format)
	if s.Blit(&rect, surface, &rect) != nil {
		panic(err)
	}
	*/

}