package main

import "testing"

func TestEval(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  int
	}{
		{"+ & *", "2+2*2", 6},
		{"- & /", "5-10/5", 3},
		{"()", "(2+2)*2", 8},
		{"^", "(2+2)*2^3", 32},
		{"!", "2+3!", 8},
		{"!", "2+3!", 8},
		{"unary -", "-2+10*(-5)", (-52)},
		{"unary +", "+5-(+10)", (-5)},
		//{"flying out", "2^(2+1)", 8},
		//{"empty", "", 0},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_, got := eval(tc.input)
			if got != tc.want {
				t.Errorf("got = %v want = %v", got, tc.want)
			}
		})
	}
}
