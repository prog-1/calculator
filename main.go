package main

import (
	"fmt"
	"math"
	"strconv"
)

// Grammar:
// expr ::= summand | summand ("+"|"-") expr
// summand ::= factor | factor ("*"|"/") summand
// factor ::= power | power ("^") factor
// power ::= (+|-) factorial | "(" expr ")"
// factorial ::= number | number ("!")

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

// factor ::= power | power ("^") factor
func factor(in []string) ([]string, int) {
	in, a := power(in)
	for {
		if len(in) == 0 || in[0] != "^" {
			return in, a
		}
		var b int
		in, b = power(in[1:])
		a = int(math.Pow(float64(a), float64(b)))
	}
}

var unaryOps = map[string]func(int) int{
	"+": func(a int) int { return a },
	"-": func(a int) int { return -a },
}

// power ::= (+|-) factorial | "(" expr ")"
func power(in []string) ([]string, int) {
	if unaryOps[in[0]] != nil { // Have unary operator
		op := unaryOps[in[0]]
		in, a := power(in[1:])
		return in, op(a)
	}
	if in[0] == "(" {
		in, x := expr(in[1:])
		if len(in) == 0 || in[0] != ")" {
			panic("want closing parentheses")
		}
		return in[1:], x
	}
	return factorial(in)
}

func factorial(in []string) ([]string, int) {
	in, a := number(in)
	if len(in) == 0 || in[0] != "!" {
		return in, a
	}
	return in[1:], int(fact(uint(a)))
}

// Returns factorial of a
func fact(a uint) uint {
	var result uint
	for i := uint(0); i <= a; i++ {
		result += i
	}
	return result
}

func number(in []string) ([]string, int) {
	x, _ := strconv.Atoi(in[0])
	return in[1:], x
}

func main() {
	fmt.Println(eval("2", "+", "2", "*", "2"))                          // 6
	fmt.Println(eval("(", "2", "+", "2", ")", "*", "2"))                // 8
	fmt.Println(eval("10", "/", "10", "*", "5"))                        // 5
	fmt.Println(eval("2", "+", "2", "^", "2"))                          // 6
	fmt.Println(eval("(", "2", "+", "2", ")", "^", "2"))                // 16
	fmt.Println(eval("-", "2"))                                         // -2
	fmt.Println(eval("2", "+", "(", "-", "2", ")"))                     // 0
	fmt.Println(eval("2", "+", "(", "-", "(", "-", "2", ")", ")"))      // 4
	fmt.Println(eval("2", "+", "(", "+", "(", "-", "2", ")", ")"))      // 0
	fmt.Println(fact(3))                                                // 1 + 2 + 3 = 6
	fmt.Println(eval("3", "!"))                                         // 6
	fmt.Println(eval("2", "+", "(", "+", "(", "-", "3", "!", ")", ")")) // -4
}
