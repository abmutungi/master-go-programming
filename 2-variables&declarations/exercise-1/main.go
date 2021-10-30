package main

func main() {
	var (
		a int
		b float64
		c bool
		d string
	)

	x, y, z := 20, 15.5, "Gopher!"

	_, _, _, _, _, _, _ = a, b, c, d, x, y, z //using blank identifier to mute the unused variable error
	// _ stays always on the left side of =

}
