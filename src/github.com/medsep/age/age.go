package main

import "fmt"

func main() {
	age := -1

	if age < 13 {
		fmt.Println("You are young!!")
	} else if age < 20 {
		fmt.Println("You are a teenager!!")
	} else if age < 30 {
		fmt.Println("You are in your twenties!!")
	} else if age < 40 {
		fmt.Println("You are in your thrities!!")
	} else if age < 50 {
		fmt.Println("You are in there!!")
	} else if age >= 50 {
		fmt.Println("You are over the hill!!")
	}
}
