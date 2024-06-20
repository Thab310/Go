package integers

import "testing"

func TestAddition(t *testing.T) {
	sum := AddInt(2, 7)
	expected := 9

	if sum != expected {
		t.Errorf("sum is %d expected %d", sum, expected)
	}
}
