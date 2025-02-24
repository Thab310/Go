package perimeter

import "testing"

func TestPerimeter(t *testing.T) {
	var want float64
	got := Perimeter(2, 3)
	want = 10.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}
