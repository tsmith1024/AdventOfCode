package main

import "testing"

func TestCount( t *testing.T) {
	cases := []struct{
		in string
		expected int
	}{
		{"(", 1},
		{"()", 0},
		{"(((", 3},
		{"(()(()(", 3},
		{"))(((((", 3},
		{"())", -1},
		{"))(", -1},
		{")))", -3},
		{")())())", -3},
	}
	for _, c := range cases {
		result := Count(c.in)
		if result != c.expected {
			t.Errorf("Count(%q) == %d, expected %d", c.in, result, c.expected)
		}
	}
}
