package day2

import "testing"

func TestWrappingArea(t *testing.T) {
	cases := []struct {
		in       []int
		expected int
	}{
		{[]int{0, 0, 0}, 0},
		{[]int{2, 3, 4}, 58},
	}
	for _, c := range cases {
		result := WrappingArea(c.in)
		if result != c.expected {
			t.Errorf("WrappingArea(%q) == %d, expected %d", c.in, result, c.expected)
		}
	}
}
