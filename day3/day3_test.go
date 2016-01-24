package main

import "testing"

func TestHouseCount(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"", 1},
		{">", 2},
		{"^>v<", 4},
		{"^v^v^v^v^v", 2},
	}
	for _, c := range cases {
		result := houseCount(c.input)
		if result != c.expected {
			t.Errorf("House count expected to be: %d\nResult is: %d", c.expected, result)
		}
	}
}

func TestNextHouse(t *testing.T) {
	cases := []struct {
		input    string
		expected house
	}{
		{"", house{0, 0}},
		{">", house{1, 0}},
		{"<", house{-1, 0}},
		{"v", house{0, -1}},
		{"^", house{0, 1}},
	}
	for _, c := range cases {
		result := nextHouse(house{0, 0}, c.input)
		if result != c.expected {
			t.Errorf("House not as expected")
		}
	}
}

func TestDualWork(t *testing.T) {
	cases := []struct {
		input    string
		expected int
	}{
		{"", 1},
		{"^v", 3},
		{"^>v<", 3},
		{"^v^v^v^v^v", 11},
	}
	for _, c := range cases {
		result := dualWork(c.input)
		if result != c.expected {
			t.Errorf("House count expected to be: %d\nResult is: %d", c.expected, result)
		}
	}
}
