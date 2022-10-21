package main

import (
	"fmt"
	"image/color"
	"os"
	"runtime"
	"time"
)

type Point struct{ X, Y int }

func (p Point) Distance(q Point) float64 {
	fmt.Println("p calculates distace to q")
	return 0.0
}

func (p *Point) ScaleBy(factor float64) {
	fmt.Println("scales p by factor")
}

type ColoredPoint struct {
	*Point
	Color color.RGBA
}

type Rocket struct{}

func (r Rocket) Launch() {
	fmt.Println("Rocket launched!")
}

func printStack() {
	var buf [4096]byte
	n := runtime.Stack(buf[:], false)
	os.Stdout.Write(buf[:n])
}

func main() {

	p := ColoredPoint{&Point{1, 1}, color.RGBA{}}
	q := ColoredPoint{&Point{2, 2}, color.RGBA{}}
	p.Distance(*q.Point)

	r := new(Rocket)
	time.AfterFunc(2*time.Second, r.Launch)

	fmt.Printf("%T\n", Point.Distance)

	time.Sleep(3 * time.Second)
}
