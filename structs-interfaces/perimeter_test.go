package structsinterfaces

import "testing"

var rectangle = Rectangle{10.0, 10.0}
const rectanglePerimeter = 40.0
const rectangleArea = 100.0



func TestPerimeter(t *testing.T) {
	got := Perimeter(rectangle)
	want := rectanglePerimeter

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	got := Area(rectangle)
	want := 100.0

	if got != want {
		t.Errorf("got %.2f want %2.f", got, want)
	}
}
