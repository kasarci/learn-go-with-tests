package structsinterfaces

import (
	"testing"
)

var rectangle = Rectangle{10.0, 10.0}
var circle = Circle{10.0}

const rectanglePerimeter = 40.0
const rectangleArea = 100.0
const circlePerimeter = 62.83185307179586
const circleArea = 314.1592653589793

func TestPerimeter(t *testing.T) {
	t.Run("rectangles", func(t *testing.T) {
		got := rectangle.Perimeter()
		want := rectanglePerimeter

		if got != want {
			t.Errorf("got %.2f want %.2f", got, want)
		}
	})

	t.Run("circles", func(t *testing.T) {
		got := circle.Perimeter()
		want := circlePerimeter

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name string
		shape Shape
		hasArea float64
	} {
		{name:"Rectangle",  shape: rectangle, hasArea: rectangleArea},
		{name:"Circle", shape: circle, hasArea: circleArea},
		{name: "Triangle", shape: Triangle{3,4}, hasArea: 6.0},
	}

	for _, sut := range areaTests {
		got := sut.shape.Area()
		if got != sut.hasArea {
			t.Errorf("%#v got %g want %g", sut.shape, got, sut.hasArea)
		}
	}
}
