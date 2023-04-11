package main

import (
	"fmt"
	"math"
)

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) *Point {
	return &Point{x, y}
}

func pointDist(point1, point2 *Point) float64 {
	return math.Abs(math.Sqrt((point1.x-point2.x)*(point1.x-point2.x) + (point1.y-point2.y)*(point1.y-point2.y)))
}
func main() {
	var p1, p2 *Point
	p1 = NewPoint(0, 0)
	p2 = NewPoint(2, 2)
	fmt.Println(pointDist(p1, p2))
}
