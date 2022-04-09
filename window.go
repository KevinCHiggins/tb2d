package tb2d
import (
	"github.com/veandco/go-sdl2/sdl"
)

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

	viewport = &sdlSurfaceWrapper{surface}

}