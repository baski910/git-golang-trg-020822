package sum1

import "testing"

func TestNums(t *testing.T) {
	tt := []struct {
		name    string
		numbers []int
		sum     int
	}{
		{"one to five", []int{1, 2, 3, 4, 5}, 15},
		{"no number", nil, 0},
		{"one and minus one", []int{1, -1}, 0},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			s := Nums(tc.numbers...)

			if s != tc.sum {
				t.Errorf("sum of %v equal to %v: %v", tc.numbers, tc.sum, s)
				//t.Fatalf("sum of %v equal to %v: %v", tc.name, tc.sum, s)
			}

		})
	}
}
