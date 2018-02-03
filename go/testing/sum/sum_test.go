package sum

import "testing"

func TestInts(t *testing.T) {
	tt := []struct {
		name   string
		nums   []int
		expSum int
	}{
		{"1 to 5", []int{1, 2, 3, 4, 5}, 15},
		{"nil", nil, 0},
		{"empty slice", []int{}, 0},
		{"1 and -1", []int{1, -1}, 0},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			gotSum := Ints(tc.nums...)
			if tc.expSum != gotSum {
				t.Fatalf("sum of %v should be %v; got %v", tc.nums, tc.expSum, gotSum)
			}
		})
	}
}
