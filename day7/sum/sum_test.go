package sum

import "testing"

func TestNums(t *testing.T) {
	s := Nums(1, 2, 3, 4, 5)

	if s != 15 {
		t.Errorf("Sum of 1 to 5 should be 15, got %d\n", s)
	}

	s = Nums()

	if s != 0 {
		t.Errorf("sum of no number should be 0, got %d\n", s)
	}

	s = Nums(1, -1)

	if s != 0 {
		t.Errorf("sum of one and minus one should be 0, got %d\n", s)
	}

}
