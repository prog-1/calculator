package main

import "testing"

func TestExpr(t *testing.T) {
	for _, tc := range []struct {
		name string
		in   string
		want int
	}{
		{"1", "-2+2*2", 2},
		{"2", "2+-2*2", -2},
		{"3", "+(2+2)*-2", -8},
		{"c4", "-(2+2)*-2", 8},
		{"5", "+(120+2*2+1)/25*10*2+100+34", 234},
		{"6", "2^10", 1024},
		{"7", "-2^10", -1024},
		{"8", "3!", 6},
		{"9", "-3!", -6},
		{"10", "(120+2*2+1)/+25*10*2+100+34+3!", 240},
		{"11", "(120+2*2+1)/25*10*2+100+34-8^2", 170},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if _, got := expr(tc.in); got != tc.want {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}
