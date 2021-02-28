package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	var a complex128
	a = cmplx.Acos(-2 + 12i)

	fmt.Println("Value of a: ", a)
}
