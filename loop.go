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
			switch t:= event.(type) {
			case *sdl.QuitEvent:
				running = false
				break
			case *sdl.MouseButtonEvent:
				if t.State == sdl.PRESSED {
					pressedClickable = getClickableContaining(int(t.X), int(t.Y))
				} else {
					releasedClickable := getClickableContaining(int(t.X), int(t.Y))
					if releasedClickable == pressedClickable && pressedClickable != nil {
						bounds := releasedClickable.GetBounds()
						releasedClickable.click(int(t.X) - bounds.X, int(t.Y) - bounds.Y)
					}
					pressedClickable = nil
				}
			}
		}
	}

}