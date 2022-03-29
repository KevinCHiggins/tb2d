package tb2d
import (
	"github.com/veandco/go-sdl2/sdl"
)

func Start(tick func()) {

	defer sdl.Quit()
	defer window.Destroy()
	running := true
	for running {
		blitDrawables()
		window.UpdateSurface()
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			}
		}
	}

}