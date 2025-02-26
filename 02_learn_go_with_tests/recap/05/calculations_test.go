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
	t.Run("Area of a rectangle", func(t *testing.T) {
		rectangle := Rectangle{10, 2}

		got := rectangle.Area()
		want := 20.0

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

	t.Run("Area of a circle", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area() //confused about assigning a method to got #Find more examples!
		want := 314.1592653589793

		if got != want {
			t.Errorf("got %g want %g", got, want)
		}
	})

}
