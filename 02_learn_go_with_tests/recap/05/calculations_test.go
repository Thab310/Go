package calculations

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10, 5}

	got := Perimeter(rectangle)
	want := 30.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	checkArea := func(t testing.TB, shape Shape, want float64) {
		t.Helper()
		got := shape.Area()
		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	}
	t.Run("Area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 2}
		checkArea(t, rectangle, 20.0)
	})

	t.Run("Area of a circle", func(t *testing.T) {
		circle := Circle{10}
		checkArea(t, circle, 314.1592653589793)
	})

}
