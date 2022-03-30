package tb2d

import "fmt"

type Rect struct {
	X, Y, W, H int
}

// needs tests
func (r Rect) contains(x, y int) bool {
	fmt.Println(r.X, x, r.W, r.Y, y, r.H)
	return (x >= r.X) && (x < r.X + r.W) && (y >= r.Y) && (y < r.Y + r.H)
}