package main

import (
	"fmt"
)

/*func main() {
	//var a int32
	const a = 15
	v := a

	fmt.Println("Address of a," & v)
}*/

func main() {
	const k = 5
	v := k
	address := &v // This is allowed
	fmt.Println("Address of a," & address)
}
