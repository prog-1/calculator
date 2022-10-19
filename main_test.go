package main

import "testing"

func TestExpresins(t *testing.T) {
	for _, tc := range []struct {
		n    string
		want int
	}{
		{"1+1", 2},
		{"1", 1},
		{"2*2", 4},
		{"10/2", 5},
		{"(2+2)*2", 8},
		{"2^5", 32},
		{"5!", 120},
		{"2+2*2", 6},
		{"3!^2", 36},
		// {"(3^2)!", 362880}, bugg
		{"10/(2+2)", 2},
		// {"2^2^2", 16}, bugg
	} {
		if _, got := eval(tc.n); got != tc.want {
			t.Errorf("npn(%v) = %v, want %v", tc.n, got, tc.want)
		}
	}
}
