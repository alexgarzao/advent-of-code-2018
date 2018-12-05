package main

type fabric struct {
	inches    [1000][1000]int
	goodClaim map[int]bool
}

func newFabric() *fabric {
	f := &fabric{}
	f.goodClaim = make(map[int]bool)

	return f
}

func (f *fabric) markInches(c claim) {
	f.goodClaim[c.id] = true
	for y := c.y; y < c.y+c.h; y++ {
		for x := c.x; x < c.x+c.w; x++ {
			currentID := f.inches[y][x]
			if currentID == 0 {
				f.inches[y][x] = c.id
			} else {
				if currentID != -1 {
					delete(f.goodClaim, currentID)
					delete(f.goodClaim, c.id)
				}
				f.inches[y][x] = -1
			}
		}
	}
}

func (f *fabric) overlaped() (result int) {
	result = 0
	for y := 0; y < 1000; y++ {
		for x := 0; x < 1000; x++ {
			if f.inches[y][x] == -1 {
				result++
			}
		}
	}
	return result
}

func (f *fabric) smallerGoodClaim() (result int) {
	result = 999999
	for key := range f.goodClaim {
		if key < result {
			result = key
		}
	}

	return result
}
