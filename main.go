package main

import (
	"bigint/bigint"
	"fmt"
)

func main() {
	a, err := bigint.NewInt("-1234567")
	if err != nil {
		panic(err)
	}

	b, err := bigint.NewInt("-000046")
	if err != nil {
		panic(err)
	}

	err = b.Set("0000464343")
	if err != nil {
		panic(err)
	}

	add := bigint.Add(a, b)
	sub := bigint.Sub(a, b)
	mul := bigint.Multiply(a, b)
	mod := bigint.Mod(a, b)
	div := bigint.Divide(a, b)
	abs := a.Abs()

	fmt.Println("a =", a)
	fmt.Println("b =", b)
	fmt.Println("added =", add)
	fmt.Println("sub =", sub)
	fmt.Println("mul =", mul)
	fmt.Println("mod =", mod)
	fmt.Println("div =", div)
	fmt.Println("abs of a =", abs)
}
