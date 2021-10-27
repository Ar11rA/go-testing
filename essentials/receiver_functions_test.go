package essentials

import "testing"

func TestSplit(t *testing.T) {
	d := Dict{}
	d = append(d, "hello")
	d = append(d, "hi")
	d = append(d, "yo")
	d = append(d, "hilo")
	d = append(d, "hiiiiiiii")
	bw, sw := d.Split(3)
	if len(bw) != 3 {
		t.Errorf("Big words should be 3 but got %v", len(bw))
	}
	if len(sw) != 2 {
		t.Errorf("Small words should be 2 but got %v", len(sw))
	}
}

func TestPrint(t *testing.T) {
	d := Dict{}
	d = append(d, "hello")
	d = append(d, "hi")
	d.Print()
}
