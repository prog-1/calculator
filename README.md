# Calculator

## Section #1: Grammar, simple implementation

### Step #1: grammar
https://go.dev/play/p/gIY8kXknAUz
### Step #2: summator
https://go.dev/play/p/V-OUJFoKFQY
### Step #3: multiplier
https://go.dev/play/p/gjyvFfySc60
### Step #4: parentheses
https://go.dev/play/p/sJ1kniqZuS4
### Step #5: left/right priority, removing recursion
https://go.dev/play/p/FyJXlyhJxpf

## Section #2: Simple string parsing

https://go.dev/play/p/4NnjIHyIqLr

## Exercises

You have to implement the following three features:
1. Power (a^b) 2^10=1024
2. Unary +/-
3. Factorial (a!) 4!=24

Grammar for (1) and (2):

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
