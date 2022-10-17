package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []string
		want  int
	}{
		{"nil", nil, 0},
		{"empty", []string{}, 0},
		{"2+2*2", []string{"2", "+", "2", "*", "2"}, 6},
		{"-(2+2!)*2", []string{"-", "(", "2", "+", "2", "!", ")", "*", "2"}, (-8)},
	} {
		t.Run(tc.name, func(t *testing.T) {
			_,got := expr(tc.input)
			if got != tc.want {
				t.Errorf("got = %v want = %v", got, tc.want)
			}
		})
	}
}
