package tb2d

import (
	"github.com/veandco/go-sdl2/sdl"
)

var viewport Graphic
var window *sdl.Window
var pxfmt *sdl.PixelFormat
var pressedClickable Clickable