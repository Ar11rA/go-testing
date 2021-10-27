package essentials

import "testing"

func TestDefaultInit(t *testing.T) {
	p := DefaultInit()
	if p.name != "default" {
		t.Errorf("Name should be defualt but got %v", p.name)
	}
}

func TestUpdatePosition(t *testing.T) {
	p := DefaultInit()
	// pointer shorthand
	// personPointer := &p
	// personPointer.UpdatePosition("SG")
	p.UpdatePosition("SG")
	if p.position != "SG" {
		t.Errorf("Pos should be SG but got %v", p.position)
	}
}
