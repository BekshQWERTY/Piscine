package main

import "github.com/01-edu/z01"

type point struct {
	x int
	y int
}

func setPoint(p *point) {
	p.x = 6 * 7
	p.y = 3 * 7
}

func main() {
	p := &point{}
	setPoint(p)

	out := []rune{
		'x', ' ', '=', ' ',
		digit(p.x / (2 * 5)),
		digit(p.x - (p.x/(2*5))*10),
		',', ' ', 'y', ' ', '=', ' ',
		digit(p.y / (3 * 3)),
		digit(p.y - (p.y/(3*3))*10),
		'\n',
	}

	for _, r := range out {
		z01.PrintRune(r)
	}
}

func digit(n int) rune {
	c := '0'
	for i := 0; i < n; i++ {
		c++
	}
	return c
}
