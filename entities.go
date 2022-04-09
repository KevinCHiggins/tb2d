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
	GetDirtyFlag() bool // bit confusing because this is a bool flag
	SetDirtyFlag()
	ClearDirtyFlag()
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
	b.ClearDirtyFlag()
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

func DeleteButton (b *Button) error {
	foundButton := false
	var newDrawables []Drawable
	for _, d := range drawables {
		if Drawable(b) != d {
			newDrawables = append(newDrawables, d)
		} else {
			foundButton = true
		}
	}
	if !foundButton {
		return errors.New("Button not found in drawables") 
	}
	foundButton = false
	var newClickables []Clickable
	for _, d := range clickables {
		if Clickable(b) != d {
			newClickables = append(newClickables, d)
		} else {
			foundButton = true
		}
	}
	if !foundButton {
		return errors.New("Button not found in clickables") 
	}
	drawables = newDrawables
	clickables = newClickables
	return nil
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

func (b *Button) SetDirtyFlag() {
	b.isClean = false
}

func (b *Button) ClearDirtyFlag() {
	b.isClean = true
}

func (b *Button) CenterInRect(holder Rect) {
	holderCenterX := holder.X + (holder.W / 2)
	holderCenterY := holder.Y + (holder.H / 2)
	origButtonBounds := b.GetBounds()
	b.SetBounds(Rect{holderCenterX - (origButtonBounds.W / 2),
		holderCenterY - (origButtonBounds.H / 2),
		origButtonBounds.W,
		origButtonBounds.H})
	debugB := b.GetBounds()
	println("It is done foist", debugB.X, debugB.Y, debugB.W, debugB.H)
	//b.SetBounds(Rect{0, 0,
	//	origButtonBounds.W,
	//	origButtonBounds.H})
	debugB = b.GetBounds()
	println("It is done", debugB.X, debugB.Y, debugB.W, debugB.H)
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
	tg.ClearDirtyFlag()
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

func (tg *TileGrid) SetDirtyFlag() {
	tg.isClean = false
}

func (tg *TileGrid) ClearDirtyFlag() {
	tg.isClean = true
}

func (tg *TileGrid) GetDirtyFlag() bool {
	return !tg.isClean
}