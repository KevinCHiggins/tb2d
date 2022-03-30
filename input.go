package tb2d

// A TileGrid is clicked if you mousedown and mouseup anywhere within it;
// meaning it is possible to start the gesture on one tile but
// end up clicking a different one

func getClickableContaining(viewportX, viewportY int) Clickable {
	for _, c := range clickables {
		if c.GetBounds().contains(viewportX, viewportY)  {
			return c
		}
	}
	println("nil")
	return nil
}