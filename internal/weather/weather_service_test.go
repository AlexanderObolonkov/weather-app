package weather

import "testing"

func TestFoo(t *testing.T) {
	foo := -2
	bar := -2
	if foo != bar {
		t.Errorf("got %d; bar %d", foo, bar)
	}
}
