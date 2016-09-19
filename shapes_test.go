package shapes_test

import (
	"github.com/paulcadman/shapes"
	"image"
	"image/color"
	"math"
	"testing"
)

func TestCircle_Area(t *testing.T) {
	var tests = []struct {
		radius   int
		expected float64
	}{
		{0, 0},
		{1, math.Pi},
		{2, math.Pi * 4},
	}

	for _, test := range tests {
		actual := shapes.NewCircle(test.radius, image.Pt(1, 1), color.Black).Area()
		if actual != test.expected {
			t.Errorf("Circle.Area() expected %f actual %f", test.expected, actual)
		}
	}
}

func TestRectangle_Area(t *testing.T) {
	var tests = []struct {
		width    int
		height   int
		expected float64
	}{
		{0, 0, 0},
		{1, 0, 0},
		{0, 1, 0},
		{1, 1, 1},
		{2, 1, 2},
	}

	for _, test := range tests {
		actual := shapes.NewRectangle(test.width, test.height, image.Pt(1, 1), color.Black).Area()
		if actual != test.expected {
			t.Errorf("Rectangle.Area() expected %f actual %f", test.expected, actual)
		}
	}
}
