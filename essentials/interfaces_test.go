package essentials

import "testing"

// Without using interfaces for mocking
func TestMeasure(t *testing.T) {
	r := Rect{width: 3, height: 4}
	c := Circle{radius: 5}
	a1, p1 := Measure(r)
	if a1 != 12 {
		t.Errorf("Area should be 12")
	}
	if p1 != 14 {
		t.Errorf("Perimeter should be 14")
	}
	a2, p2 := Measure(c)
	if a2 != 78.53981633974483 {
		t.Errorf("Area should be 78.53981633974483")
	}
	if p2 != 31.41592653589793 {
		t.Errorf("Perimeter should be 31.41592653589793")
	}
}

type mockGeometry struct{}

func (g mockGeometry) area() float64 {
	return 10.0
}

func (g mockGeometry) perim() float64 {
	return 20.0
}

// Using interfaces for mocking
func TestMeasureWithMocks(t *testing.T) {
	r := mockGeometry{}
	a1, p1 := Measure(r)
	if a1 != 10 {
		t.Errorf("Area should be 10")
	}
	if p1 != 20 {
		t.Errorf("Perimeter should be 20")
	}
}
