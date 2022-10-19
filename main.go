package main

import (
	"fmt"
	"strconv"
)

func expr(in string) (string, int) { return sum(in) }

var addOps = map[byte]func(int, int) int{
	'+': func(a, b int) int { return a + b },
	'-': func(a, b int) int { return a - b },
}

var mulOps = map[byte]func(int, int) int{
	'*': func(a, b int) int { return a * b },
	'/': func(a, b int) int { return a / b },
	'^': func(a, b int) int {
		if b < 0 {
			panic("negative power is not supported")
		}
		x := 1
		for ; b > 0; b-- {
			x *= a
		}
		return x
	},
}

var unaryOps = map[byte]func(int) int{
	'+': func(a int) int { return a },
	'-': func(a int) int { return -a },
}

var factOp = map[byte]func(int) int{
	'!': func(a int) int {
		x := a
		for ; a > 0; a-- {
			x *= a
		}
		return x
	},
}

func sum(in string) (string, int) {
	in, a := mul(in)
	if in == "" || addOps[in[0]] == nil {
		return in, a
	}
	op := addOps[in[0]]
	in, b := sum(in[1:])
	return in, op(a, b)
}

func mul(in string) (string, int) {
	in, a := power(in)
	for {
		if len(in) == 0 || mulOps[in[0]] == nil {
			return in, a
		}
		op := mulOps[in[0]]
		var b int
		in, b = power(in[1:])
		a = op(a, b)
	}
}

func unary(in string) (string, int) {
	in, a := power(in)
	if len(in) == 0 || unaryOps[in[0]] == nil {
		return in, a
	}
	op := unaryOps[in[0]]
	in, b := power(in[1:])
	return in, op(b)
}

func power(in string) (string, int) {
	in, a := fact(in)
	if len(in) == 0 || in[0] != '^' {
		return in, a
	}
	op := mulOps[in[0]]
	in, b := fact(in[1:])
	a = op(a, b)
	return in, a
}

func fact(in string) (string, int) {
	in, a := value(in)
	if len(in) == 0 || in[0] == '!' {
		return in, a
	}
	op := factOp[in[0]]
	a = op(a)
	return in[1:], a
}

func value(in string) (string, int) {
	if in[0] == '(' {
		in, x := sum(in[1:])
		if in == "" || in[0] != ')' {
			panic("want closing parentheses")
		}
		return in[1:], x
	}
	return num(in)
}

func num(in string) (string, int) {
	fmt.Println(in)
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

// expr  ::= sum
// sum   ::= mul | mul [+-] sum
// mul   ::= unary | unary [*/] mul
// unary ::= power | [+-] power
// power ::= fact | fact ^ power
// fact  ::= value | value !
// value ::= ( expr ) | num
// num   ::= digit | digit num
// digit ::= ...

func main() {
	fmt.Println(expr("2+2*2"))
	fmt.Println(expr("2*2+2"))
	fmt.Println(expr("(2+2)*2"))
	fmt.Println(expr("2*(2+2)"))
	fmt.Println(expr("(120+2*2+1)/25*10*2+100+34"))
}
