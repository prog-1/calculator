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
4. Find and fix a bug. Yes, we have it!

> **Note**
> Unary `-` priority is lower than power or factorial. It means that `-2^4` is `-(2^4)` and `-3!` is `-(3!)`

### Grammar

> **Note**
> This grammar is operator priority based. It is not very natural to read. Check another version of the grammar below.
> Feel free to use any version that feels better to you.

```
Association:
  [-] N/A
  [L] Left
  [R] Right
  [*] Left or right

expr      [*] ::= summand | summand ("+"|"-") expr
summand   [L] ::= factor | factor ("*"|"/") summand
factor    [-] ::= ("+"|"-") unary | unary
unary     [R] ::= power | power "^" unary
power     [-] ::= factorial | factorial "!"
factorial [-] ::= "(" expr ")" | number
number    [*] ::= digit | digit number
digit     [-] ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
```

> **Note**
> Alternative version of the grammar is operator focused. Rules are named differently, but they are operation centric - `sum`
> describes how to sum terms, `mul` describes how to multiple factors, etc. Functionally these two grammars are the same.

```bnf
Corresponding associations apply!

<expr>  ::= <sum>
<sum>   ::= <mul> | <mul> ("+" | "-") <sum>
<mul>   ::= <unary> | <unary> ("*" | "/") <mul>
<unary> ::= ("+" | "-") <power> | <power>
<power> ::= <fact> | <fact> "^" <power>
<fact>  ::= <value> | <value> "!"
<value> ::= "(" <expr> ")" | <num>
<num>   ::= <digit> | <digit> <num>
<digit> ::= "0" | "1" | "2" | "3" | "4" | "5" | "6" | "7" | "8" | "9"
```

You can play with the matcher for the grammar [here](https://bnfplayground.pauliankline.com/?bnf=%3Cexpr%3E%20%20%3A%3A%3D%20%3Csum%3E%0A%3Csum%3E%20%20%20%3A%3A%3D%20%3Cmul%3E%20%7C%20%3Cmul%3E%20%28%22%2B%22%20%7C%20%22-%22%29%20%3Csum%3E%0A%3Cmul%3E%20%20%20%3A%3A%3D%20%3Cunary%3E%20%7C%20%3Cunary%3E%20%28%22%2a%22%20%7C%20%22%2F%22%29%20%3Cmul%3E%0A%3Cunary%3E%20%3A%3A%3D%20%28%22%2B%22%20%7C%20%22-%22%29%20%3Cpower%3E%20%7C%20%3Cpower%3E%0A%3Cpower%3E%20%3A%3A%3D%20%3Cfact%3E%20%7C%20%3Cfact%3E%20%22%5E%22%20%3Cpower%3E%0A%3Cfact%3E%20%20%3A%3A%3D%20%3Cvalue%3E%20%7C%20%3Cvalue%3E%20%22%21%22%0A%3Cvalue%3E%20%3A%3A%3D%20%22%28%22%20%3Cexpr%3E%20%22%29%22%20%7C%20%3Cnum%3E%0A%3Cnum%3E%20%20%20%3A%3A%3D%20%3Cdigit%3E%20%7C%20%3Cdigit%3E%20%3Cnum%3E%0A%3Cdigit%3E%20%3A%3A%3D%20%220%22%20%7C%20%221%22%20%7C%20%222%22%20%7C%20%223%22%20%7C%20%224%22%20%7C%20%225%22%20%7C%20%226%22%20%7C%20%227%22%20%7C%20%228%22%20%7C%20%229%22).

