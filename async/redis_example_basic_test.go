package async

import "testing"

func TestClientSetGet(t *testing.T) {
	val := ClientSetGet()
	if val != "value" {
		t.Errorf("val should be value")
	}
}
