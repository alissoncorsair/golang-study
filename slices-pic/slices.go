package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	a := make([][]uint8, dy)

	//i range a
	for i := range a {
		a[i] = make([]uint8, dx)
		for b := range a[i] {
			a[i][b] = uint8(i * b)
		}
	}
	return a
}

func main() {
	pic.Show(Pic)
}
