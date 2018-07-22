package main

import "fmt"

type Point struct {
	x, y float64
}

func (p Point) Add(q Point) Point {
	return Point{p.x + q.x, p.y + q.y}
}

func (p Point) Sub(q Point) Point {
	return Point{p.x - q.x, p.y - q.y}
}

type Path []Point

func (path Path) TranslateBy(offset Point, add bool) {
	// TODO: this?
	var op func(p, q Point) Point

	if add {
		op = Point.Add
	} else {
		op = Point.Sub
	}

	for i := range path {
		path[i] = op(path[i], offset)
	}
}

func main() {
	var path Path

	path = append(path, Point{1, 1})
	path = append(path, Point{2, 2})
	path = append(path, Point{3, 3})

	offSet := Point{1, 1}

	//path.TranslateBy(offSet, true)
	path.TranslateBy(offSet, false)

	for _, v := range path {
		fmt.Println(v.x, " ", v.y)
	}
}
