package main

import (
	"fmt"
	"time"
)

func main() {

	//i := 10

	for i := 10; i > 0; i-- {
		fmt.Println(i)
		time.Sleep(time.Second)
		//i -= 1
	}
	fmt.Println("Happy new year!")
}
