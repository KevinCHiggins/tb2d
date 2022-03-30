package tb2d

import ("fmt"
	"errors"
)



type Entity interface {
	GetBounds() Rect
	SetBounds(r Rect)
}

type Drawable interface {
	Entity
	GetDirtyFlag() bool
	SetDirtyFlag(bool)
	draw() Graphic
}

type Clickable interface {
	Entity
	click(clientX, clientY int)
}
var drawables []Drawable

var clickables []Clickable

type Button struct {
	bounds Rect
	graphic Graphic
	onclick func()
	isClean bool // rather than isDirty so it starts dirty (zero-value is dirty)
}

type TileGrid struct {
	bounds Rect
	tileGraphics []Graphic
	columnsAmount, rowsAmount int
	grid []int // indices into tileGraphics, as descending rows arranged contiguously
	onclick func(int, int)
	isClean bool
}

func (b *Button) draw() Graphic {
	b.SetDirtyFlag(false)
	println("Drew button")
	return b.graphic
}

func NewButtonFromFile(filePath string, buttonClick func(), viewportX, viewportY int) *Button {
	b := Button{Rect{}, NewGraphicFromFile(filePath), buttonClick, false}
	b.bounds = Rect{viewportX, viewportY, b.graphic.GetWidth(), b.graphic.GetHeight()}
	drawables = append(drawables, Drawable(&b))
	clickables = append(clickables, Clickable(&b))
	return &b
}

func (b *Button) click(clientX, clientY int) {
	b.onclick()
}

func (b *Button) GetBounds() Rect {
	return b.bounds
}

func (b *Button) SetBounds(r Rect) {
	b.bounds = r
}

func (b *Button) GetDirtyFlag() bool {
	return !b.isClean
}

func (b *Button) SetDirtyFlag(isDirty bool) {
	b.isClean = !isDirty
}

func (tg *TileGrid) draw() Graphic {
	graphicBounds := tg.GetBounds()
	tileWidth :=  tg.tileGraphics[0].GetWidth()
	tileHeight := tg.tileGraphics[0].GetHeight()
	entireGrid := NewBlankGraphic(graphicBounds.W, graphicBounds.H)
	for i, _ := range tg.grid {
		blitY := (i / tg.rowsAmount) * tileHeight
		blitX := (i  % tg.rowsAmount) * tileWidth
		fmt.Println("Blit coords:", blitX, blitY)
		j := tg.grid[i]
		tg.tileGraphics[j].Blit(entireGrid, Rect{blitX, blitY, tileWidth, tileHeight})
	}
	tg.SetDirtyFlag(false)
	println("Drew tilegrid")
	return entireGrid
}

func (tg *TileGrid) GetGrid() []int {
	return tg.grid
}

func (tg *TileGrid) GetTileAt(x, y int) int {
	return tg.grid[x + y * tg.columnsAmount]
}

func (tg *TileGrid) SetTileAt(x, y, i int) {
	tg.grid[x + y * tg.columnsAmount] = i
}

func NewTileGridFromFiles(filePaths []string, gridClick func(int, int), columnsAmount, rowsAmount, viewportX, viewportY int) *TileGrid {
	if len(filePaths) == 0 {
		panic(errors.New("Empty slice of tile file paths provided - can't load any tiles"))
	}
	tileGraphics := make([]Graphic, len(filePaths))
	for i, filePath := range filePaths {
		tileGraphics[i] = NewGraphicFromFile(filePath)
		if i > 0 {
			if tileGraphics[i].GetWidth() != tileGraphics[i - 1].GetWidth() || tileGraphics[i].GetHeight() != tileGraphics[i - 1].GetHeight() {
				panic(errors.New(fmt.Sprintf("Tile", i, " does not have same dimensions as previous tile")))
			}
		}
	}
	tg := TileGrid{Rect{}, tileGraphics, columnsAmount, rowsAmount, make([]int, columnsAmount * rowsAmount), gridClick, false}

	// we have checked that all tile Graphics have the same dimensions, so we can use [0] here
	tg.bounds = Rect{viewportX, viewportY, tg.tileGraphics[0].GetWidth() * columnsAmount, tg.tileGraphics[0].GetHeight() * rowsAmount}
	drawables = append(drawables, Drawable(&tg))
	clickables = append(clickables, Clickable(&tg))
	return &tg
}

func (tg *TileGrid) click(clientX, clientY int) {
	tg.onclick(clientX / (tg.GetBounds().W / tg.columnsAmount), clientY / (tg.GetBounds().H / tg.rowsAmount))
}

func (tg *TileGrid) GetBounds() Rect {
	return tg.bounds
}

func (tg *TileGrid) SetBounds(r Rect) {
	tg.bounds = r
}

func (tg *TileGrid) SetDirtyFlag(isDirty bool) {
	tg.isClean = !isDirty
}

func (tg *TileGrid) GetDirtyFlag() bool {
	return !tg.isClean
}

/*
type TileGrid struct {
	tiles []*Graphic
}

func (t *TileGrid) draw() *Graphic {
	// make new Graphic (needs a function to do that) and blit tiles into it
	// what about bounds checks (on creation as well as drawing?)
}
func (t *TileGrid) click(clientX, clientY int) {

}
*/
