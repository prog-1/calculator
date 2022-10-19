package main

import (
	"fmt"
	"strconv"
)

func eval(in string) (result int) {
	_, result = expr(in)
	return
}

var addOps = map[byte]func(int, int) int{
	'+': func(a, b int) int { return a + b },
	'-': func(a, b int) int { return a - b },
}

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

var unaryOps = map[byte]func(int) int{
	'+': func(a int) int { return a },
	'-': func(a int) int { return -a },
}

func factor(in string) (string, int) {
	if in == "" || unaryOps[in[0]] == nil {
		return unary(in)
	}
	op := unaryOps[in[0]]
	in, a := unary(in[1:])
	return in, op(a)
}

func unary(in string) (string, int) {
	in, a := power(in)
	if in == "" || in[0] != '^' {
		return in, a
	}
	result := 1
	in, b := power(in[1:])
	for ; b > 0; b-- {
		result *= a
	}
	return in, result
}

func power(in string) (string, int) {
	in, a := factorial(in)
	if in == "" || in[0] != '!' {
		return in, a
	}
	if a < 0 {
		panic("negative number is not supported")
	}
	result := 1
	for ; a > 0; a-- {
		result *= a
	}
	return in[1:], result
}

func factorial(in string) (string, int) {
	if in[0] == '(' {
		in, x := expr(in[1:])
		if in == "" || in[0] != ')' {
			panic("want closing parentheses")
		}
		return in[1:], x
	}
	return number(in)
}

// expr      [*] ::= summand | summand ("+"|"-") expr
// summand   [L] ::= factor | factor ("*"|"/") summand
// factor    [-] ::= ("+"|"-") unary | unary
// unary     [R] ::= power | power "^" unary
// power     [-] ::= factorial | factorial "!"
// factorial [-] ::= "(" expr ")" | number
// number    [*] ::= digit | digit number
// digit     [-] ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"

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

func main() {
	fmt.Println(eval("2+2*2"))
	fmt.Println(eval("2*2+2"))
	fmt.Println(eval("(2+2)*2"))
	fmt.Println(eval("2*(2+2)"))
	fmt.Println(eval("(120+2*2+1)/25*10*2+100+34"))
	fmt.Println(eval("3^3"))
	fmt.Println(eval("3!"))
	fmt.Println(eval("-5+3"))
	fmt.Println(eval("-(2+2)"))
	fmt.Println(eval("-(2+2)*2"))
}
