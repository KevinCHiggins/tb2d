package tb2d


// put in window.go?
// draw

func blitDrawables() {
	
	for _, d := range drawables {
		if d.GetDirtyFlag() {
			g := d.draw()
			g.Blit(viewport, d.GetBounds())			
		}
	}
	
}