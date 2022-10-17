# Calculator

## Section #1: Grammar basics. Simple infix calculator.

In this section the user input is represented as a slice of strings, where each iteam is a token. Because of that there is no need to parse anything.

> **Note**
> We assume input is always valid, so no extensive error checking is made.

1. Evaluating numbers.

   https://go.dev/play/p/gIY8kXknAUz
   
2. Evaluating additions.

   https://go.dev/play/p/V-OUJFoKFQY

3. Operator priorities. Evaluating multiplication.

   https://go.dev/play/p/gjyvFfySc60

4. Evaluating parentheses.

   https://go.dev/play/p/sJ1kniqZuS4

5. Left/right order. Non-recursive left order implementation.

   https://go.dev/play/p/FyJXlyhJxpf

Now we can evaluate simple expressions e.g.

```go
fmt.Println(eval("2", "+", "2", "*", "2"))
fmt.Println(eval("(", "2", "+", "2", ")", "*", "2"))
fmt.Println(eval("10", "/", "10", "*", "5"))
// Output:
// [] 6
// [] 8
// [] 5
```

## Section #2: Simple string parsing

In this section we represent the user input as a string.

> **Info**
> We also assume the input is correct. No extensive error checking.

https://go.dev/play/p/4NnjIHyIqLr

```go
fmt.Println(eval("(120+2*2+1)/25*10*2+100+34"))
// Output: 234
```

## Exercises

### Tasks

> **Note**
> Please solve the tasks using the `string` evaluation version (from Section #2), not the `[]string` one.

> **Note**
> Feel free to join all of the features in a single program. You don't have to implement three different ones!

You have to implement the following three features:

1. Power `a^b` e.g. `2^10` should return `1024`
2. Factorial `a!` e.g. `4!` is `24`.
3. Unary `+`/`-` for expressions like `-2*5`, `5*-2`, `-(2+2)*2`, etc.

> **Note**
> Unary `-` priority is lower than power or factorial. It means that `-2^4` is `-(2^4)` and `-3!` is `-(3!)`

### Grammar

```
// Rule order:
// [-] no order
// [L] left
// [R] right
// [*] left or right

// expr      [*] ::= summand | summand ("+"|"-") expr
// summand   [L] ::= factor | factor ("*"|"/") summand
// factor    [-] ::= ("+"|"-") unary | unary
// unary     [R] ::= power | power "^" unary
// power     [-] ::= factorial | factorial "!"
// factorial [-] ::= "(" expr ")" | number
// number    [*] ::= digit | digit number
// digit     [-] ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
```
