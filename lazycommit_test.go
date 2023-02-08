package lazycommit

import "testing"

func TestAdd2(t *testing.T) {
	if a := Add2(1); a != 3 {
		t.Fatalf("Add2(1) = %d; want 3", a)
	}
}
