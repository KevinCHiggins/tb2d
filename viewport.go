package tb2d

// draw

func blitDrawables() {
	
	for _, d := range drawables {
		if d.GetDirtyFlag() {
			g := d.draw()
			g.Blit(viewport, d.GetBounds())			
		}
	}
	
}