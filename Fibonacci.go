package main

func fibonacci() func() int {
	var x, y, z = 0, 1, 0
	return func() int {
		z = x
		x = y
		y = x + z
		return z
	}
}

