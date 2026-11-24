package main

import "testing"

func TestIsHappy(t *testing.T) {
	cases := []struct {
		in   int
		want bool
	}{
		{19, true}, // 82 -> 68 -> 100 -> 1
		{2, false},
		{1, true},
		{7, true},
		{4, false}, // the heart of the unhappy cycle
		{100, true},
		{1111111, true}, // seven 1s -> 7 -> happy
	}
	for _, c := range cases {
		if got := IsHappy(c.in); got != c.want {
			t.Errorf("IsHappy(%d) = %v, want %v", c.in, got, c.want)
		}
	}
}
