package main

import (
	"math"
	"strconv"
)

func eval(in string) (string, int) {
	in, a := expr(in)

	// if len(in) != 0 {
	// 	tmp := strconv.Itoa(a)
	// 	return factorial(tmp + in)
	// }
	return in, a
}

var addOps = map[byte]func(int, int) int{
	'+': func(a, b int) int { return a + b },
	'-': func(a, b int) int { return a - b },
}

// expr ::= summand | summand ("+"|"-") expr
func expr(in string) (string, int) {
	in, a := summand(in)
	if in == "" || addOps[in[0]] == nil {
		return in, a
	}
	op := addOps[in[0]]
	in, b := expr(in[1:])
	return in, op(a, b)
}

var mulOps = map[byte]func(int, int) int{
	'*': func(a, b int) int { return a * b },
	'/': func(a, b int) int { return a / b },
}

// summand ::= factor | factor ("*"|"/") summand
func summand(in string) (string, int) {
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

// factor ::= "(" expr ")" | number
func factor(in string) (string, int) {
	if in[0] == '(' {
		in, x := expr(in[1:])
		if len(in) == 0 || in[0] != ')' {
			panic("want closing parentheses")
		}
		return in[1:], x
	}
	return power(in)
}
func power(in string) (string, int) {
	in, a := factorial(in)
	if len(in) == 0 || in[0] != '^' {
		return in, a
	}
	in, pow := factorial(in[1:])
	return in, int(math.Pow(float64(a), float64(pow)))
}
func factorial(in string) (string, int) {
	in, a := number(in)
	if len(in) == 0 || in[0] != '!' {
		return in, a
	}
	f := 1
	for i := 1; i < a+1; i++ {
		f *= i
	}
	return in[1:], f
}

func number(in string) (string, int) {
	n := 0
	for n < len(in) && '0' <= in[n] && in[n] <= '9' {
		n++
	}
	if n == 0 {
		panic("want a number")
	}
	x, _ := strconv.Atoi(in[:n])
	return in[n:], x
}
