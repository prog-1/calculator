package main

import "testing"

func TestExpresins(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input string
		want  int
	}{
		{"simple sum", "1+1", 2},
		{"simple mult", "1*2", 2},
		{"correct operation order", "2*2+2", 6},
		{"incorrect operation order", "2+2*2", 6},
		{"Brackets", "(2+2)*2", 8},
		{"Power", "2^2", 4},
		{"factorial", "4!", 24},
	} {
		if _, got := eval(tc.input); got != tc.want {
			t.Errorf("%v: Wrong answer on problem %v, got %v, want %v.", tc.name, tc.input, got, tc.want)
		}
	}
}
