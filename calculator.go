package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	fmt.Println(eval("3^2"))
	fmt.Println(eval("2*2+2"))
	fmt.Println(eval("(2+2)*2"))
	fmt.Println(eval("2*(2+2)"))
	fmt.Println(eval("(120+2*2+1)/25*10*2+100+34"))
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

//expr ::= summand | summand ("+"|"-") expr
func expr(in string) (string, int) {
	in, a := summand(in)
	if in == "" || addOps[in[0]] == nil {
		return in, a
	}
	op := addOps[in[0]]
	in, b := expr(in[1:])
	return in, op(a, b)
}

//summand ::= factor | factor ("*"|"/")
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

// factor ::= power | "(" expr ")"
func factor(in string) (string, int) {
	if in[0] == '(' {
		in, x := expr(in[1:])
		if len(in) == 0 || in[0] != ')' {
			panic("wanna closing parentheses")
		}
		return in[1:], x
	}
	return power(in)
}

// power ::= number | number (^)
func power(in string) (string, int) {
	in, a := number(in)
	for {
		if len(in) == 0 || powOps[in[0]] == nil {
			return in, a
		}
		op := powOps[in[0]]
		var b int
		in, b = number(in[1:])
		a = op(a, b)
	}
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
