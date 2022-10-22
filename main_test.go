package main

import (
	"testing"
)

func TestEval(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want int
	}{
		{"-2+2*2", 2},
		{"2+-2*2", -2},
		{"+(2+2)*-2", -8},
		{"-(2+2)*-2", 8},
		{"+(120+2*2+1)/25*10*2+100+34", 234},
		{"2^10", 1024},
		{"3!", 6},
	} {
		if _, got := eval(tc.in); got != tc.want {
			t.Errorf("eval(%v) = %v, want = %v", tc.in, got, tc.want)
		}
	}
}
