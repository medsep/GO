package main

import (
	"fmt"
)

func main() {
	var a int32
	//const a = 15
	a = 15
	fmt.Println("Address of a," & a)
}

/*
func main() {
	const k = 5
	v := k
	address := &v // This is allowed
	fmt.Println("Address of a," & address)
}*/
