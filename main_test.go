package main

import "testing"

func TestEval(t *testing.T) {
	for _, tc := range []struct {
		input string
		want  int
	}{
		{"5*-2", -10},
		{"4!", 24},
		{"2*2^3", 16},
		{"1+2^3*2-27", -10},
		{"2+2*2", 6},
		{"2*2+2", 6},
		{"(2+2)*2", 8},
		{"10/5*2", 4},
		{"2*(2+2)", 8},
		{"(120+2*2+1)/25*10*2+100+34", 234},
	} {
		if _, got := eval(tc.input); got != tc.want {
			t.Errorf("eval:%v, got %v, want %v.", tc.input, got, tc.want)
		}
	}
}
