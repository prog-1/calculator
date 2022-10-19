package main

import (
	"fmt"
	"math"
	"strconv"
)

func eval(in ...string) ([]string, int) { return expr(in) }

var addOps = map[string]func(int, int) int{
	"+": func(a, b int) int { return a + b },
	"-": func(a, b int) int { return a - b },
}

// expr ::= summand | summand ("+"|"-") expr
func expr(in []string) ([]string, int) {
	in, a := summand(in)
	if len(in) == 0 || addOps[in[0]] == nil {
		return in, a
	}
	op := addOps[in[0]]
	in, b := expr(in[1:])
	return in, op(a, b)
}

var mulOps = map[string]func(int, int) int{
	"*": func(a, b int) int { return a * b },
	"/": func(a, b int) int { return a / b },
}

// summand ::= factor | factor ("*"|"/") summand
func summand(in []string) ([]string, int) {
	in, a := factor(in)
	for {
		if len(in) == 0 || mulOps[in[0]] == nil {
			return in, a
		}
		op := mulOps[in[0]]
		var b int
		in, b = factor(in[1:])
		a = op(a, b)
	}
}

// factor ::= power | "(" expr ")"
func factor(in []string) ([]string, int) {
	if in[0] == "(" {
		in, x := expr(in[1:])
		if len(in) == 0 || in[0] != ")" {
			panic("want closing parentheses")
		}
		return in[1:], x
	}
	return power(in)
}

// power ::= number ^ number | number
func power(in []string) ([]string, int) {
	in, x := number(in)
	if len(in) == 0 || in[0] != "^" {
		return in, x
	}
	in, y := number(in[1:])
	return in, int(math.Pow(float64(x), float64(y)))
}

func number(in []string) ([]string, int) {
	x, _ := strconv.Atoi(in[0])
	return in[1:], x
}

func main() {
	fmt.Println(eval("2", "+", "2", "*", "2"))
	fmt.Println(eval("(", "2", "+", "2", ")", "*", "2"))
	fmt.Println(eval("10", "/", "10", "*", "5"))
	fmt.Println(eval("10", "^", "10", "*", "3", "^", "2"))

}
