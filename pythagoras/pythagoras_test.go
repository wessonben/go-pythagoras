package pythagoras

import (
	"testing"
)

var tests = []struct {
	a          float64
	b          float64
	hypotenuse float64
	area       float64
	perimeter  float64
}{
	{-3.0, -4.0, 0.0, 0.0, 0.0},
	{0.0, 0.0, 0.0, 0.0, 0.0},
	{3.0, 4.0, 5.0, 6.0, 12.0},
	{30.0, 40.0, 50.0, 600.0, 120.0},
	{300.0, 400.0, 500.0, 60000.0, 1200.0},
}

func TestGetHypotenuse(t *testing.T) {
	for _, test := range tests {
		expected := test.hypotenuse
		actual := GetHypotenuse(test.a, test.b)
		if actual != expected {
			t.Errorf("Hypotenuse(%f,%f): expected %f, actual %f", test.a, test.b, test.hypotenuse, actual)
		}
	}
}

func TestGetArea(t *testing.T) {
	for _, test := range tests {
		expected := test.area
		actual := GetArea(test.a, test.b)
		if actual != expected {
			t.Errorf("Area(%f,%f): expected %f, actual %f", test.a, test.b, test.area, actual)
		}
	}
}

func TestGetPerimeter(t *testing.T) {
	for _, test := range tests {
		expected := test.perimeter
		actual := GetPerimeter(test.a, test.b)
		if actual != expected {
			t.Errorf("Area(%f,%f): expected %f, actual %f", test.a, test.b, test.perimeter, actual)
		}
	}
}
