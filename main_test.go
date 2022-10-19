package main

import "testing"

func TestEval(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   string
		want int
	}{
		{"case-1", "-2+2*2", 2},
		{"case-2", "2+-2*2", -2},
		{"case-3", "+(2+2)*-2", -8},
		{"case-4", "-(2+2)*-2", 8},
		{"case-5", "+(120+2*2+1)/25*10*2+100+34", 234},
		{"case-6", "2^10", 1024},
		{"case-7", "-2^10", -1024},
		{"case-8", "3!", 6},
		{"case-9", "-3!", -6},
		{"case-10", "(120+2*2+1)/+25*10*2+100+34+3!", 240},
		{"case-11", "(120+2*2+1)/25*10*2+100+34-8^2", 170},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, got := eval(tc.in)
			if got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
