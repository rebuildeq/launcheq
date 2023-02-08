package client

import "testing"

func TestPatch(t *testing.T) {
	c, err := New("", "", "")
	if err != nil {
		t.Fatalf("new: %s", err)
	}
	c.Patch()
}
