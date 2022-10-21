package main

import "testing"

func TestExpresins(t *testing.T) {
	for _, tc := range []struct {
		name    string
		problem []string
		want    int
	}{
		{"simple summation", []string{"2", "+", "2"}, 4},
		{"simple multiping", []string{"2", "*", "3"}, 6},
		{"Operation oreder is correct", []string{"2", "*", "2", "+", "2"}, 6},
		{"Operation oreder is not correct", []string{"2", "+", "2", "*", "2"}, 6},
		{"Brackets", []string{"(", "2", "+", "2", ")", "*", "2"}, 8},
		{"Power", []string{"(", "2", "^", "6", ")", "*", "2"}, 128},
		{"factorial", []string{"(", "4", "!", "^", "6", ")", "*", "2"}, 382205952},
	} {
		if _, got := eval(tc.problem...); got != tc.want {
			t.Errorf("%v: Wrong answer on problem %v, got %v, want %v.", tc.name, tc.problem, got, tc.want)
		}
	}
}
