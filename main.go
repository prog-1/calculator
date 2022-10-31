package main

import (
	"fmt"
	"math"
	"strconv"
)

func eval(in string) (string, int) { return expr(in) }

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
		if in == "" || in[0] != ')' {
			panic("want closing parentheses")
		}
		return in[1:], x
	}
	return number(in)
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

// power ::= factorial | factorial "!"
func power(in string) (string, int) {
	in, a := factorial(in)
	if in == "" || in[0] != '!' {
		return in, a
	}
	if a < 0 {
		panic("negative number")
	}
	result := 1
	for ; a > 0; a-- {
		result *= a
	}
	return in[1:], result
}
func unary(in string) (string, int) {
	in, a := power(in)
	for {
		if len(in) == 0 || in[0] != '^' {
			return in, a
		}
		var b int
		in, b = power(in[1:])
		a = int(math.Pow(float64(a), float64(b)))
	}
}

// factorial ::= "(" expr ")" | number
func factorial(in string) (string, int) {
	if in[0] == '(' {
		in, x := expr(in[1:])
		if in == "" || in[0] != ')' {
			panic("want correct parentheses")
		}
		return in[1:], x
	}
	return number(in)
}

func main() {
	fmt.Println(eval("2+2*2"))
	fmt.Println(eval("2*2+2"))
	fmt.Println(eval("(2+2)*2"))
	fmt.Println(eval("2*(2+2)"))
	fmt.Println(eval("(120+2*2+1)/25*10*2+100+34"))
}
