package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(eval("-2*5"))
	// fmt.Println(eval("2*2+2"))
	// fmt.Println(eval("(2+2)*2"))
	// fmt.Println(eval("2*(2+2)"))
	// fmt.Println(eval("(120+2*2+1)/25*10*2+100^2+34^10"))
}

func eval(in string) (string, int) { return expr(in) } // in ...string - we can add unfixed amount of strings and treat this input as slice

var addOps = map[byte]func(int, int) int{
	'+': func(a, b int) int { return a + b },
	'-': func(a, b int) int { return a - b },
}

var mulOps = map[byte]func(int, int) int{
	'*': func(a, b int) int { return a * b },
	'/': func(a, b int) int { return a / b },
}

var powOps = map[byte]func(int, int) int{
	'^': func(a, b int) int { return int(math.Pow(float64(a), float64(b))) },
}

var unaryOps = map[byte]func(int) int{
	'+': func(a int) int { return a },
	'-': func(a int) int { return -a },
}

//expr ::= summand | summand ('+'|'-') expr
func expr(in string) (string, int) {
	in, a := summand(in)
	if in == "" || addOps[in[0]] == nil {
		return in, a
	}
	op := addOps[in[0]]
	in, b := expr(in[1:])
	return in, op(a, b)
}

//summand ::= factor | factor ('*'|'/')
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

// factor ::= unary | '(' expr ')'
func factor(in string) (string, int) {
	if in[0] == '(' {
		in, x := expr(in[1:])
		if len(in) == 0 || in[0] != ')' {
			panic("wanna closing parentheses")
		}
		return in[1:], x
	}
	return unary(in)
}

// unary ::= power | ('+'|'-') power
func unary(in string) (string, int) {

	if len(in) == 0 || unaryOps[in[0]] == nil {
		in, a := power(in)
		return in, a
	}
	op := unaryOps[in[0]]
	var b int
	in, b = power(in[1:])
	return in, op(b)

	// if in[0] == '+' {
	// 	in = in[1:]
	// 	in, n = power(in)
	// } else if in[0] == '-' {
	// 	in = in[1:]
	// 	in, n = power(in)
	// 	n = -n
	// } else {
	// 	in, n = power(in)
	// }
	// in, n = power(in)
	// return in, n

}

// power ::= factorial | factorial ('^')
func power(in string) (string, int) {

	in, a := factorial(in)
	for {
		if len(in) == 0 || powOps[in[0]] == nil {
			return in, a
		}
		op := powOps[in[0]]
		var b int
		in, b = factorial(in[1:])
		a = op(a, b)
	}
}

// factorial ::= number | number ('!')
func factorial(in string) (string, int) {
	in, a := number(in)
	if len(in) == 0 || in[0] != '!' {
		return in, a
	}
	in = in[1:]
	res := 1
	for i := a; i > 0; i-- {
		res *= i
	}
	return in, res
}

//number.. just a number
func number(in string) (string, int) {
	n := 0
	for n < len(in) && '0' <= in[n] && in[n] <= '9' {
		n++
	}
	if n == 0 {
		panic("wanna number")
	}
	x, _ := strconv.Atoi(in[:n])
	return in[n:], x
}

//expr ::= summand | summand ("+"|"-") expr
//summand ::= factor | factor ("*"|"/")
//factor ::= unary | "(" expr ")"
//unary ::= power | ('+'|'-') power
//power ::= factorial | factorial ("^")
//factorial ::= number | number ("!")
//number ::= number
