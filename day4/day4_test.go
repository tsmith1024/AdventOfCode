package main

import "testing"

func TestMD5Search(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"abcdef", 609043},
		{"pqrstuv", 1048970},
	}

	for _, c := range cases {
		result := findHashExtension(c.input, "00000")
		if result != c.expected {
			t.Errorf("Hash extension expected to be: %d\nResult is: %d\n", c.expected, result)
		}
	}
}
