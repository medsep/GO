package main

import "fmt"

func max(i int, j int, k int) {
	fmt.Println("address of k:", &k)
	if i > j {
		//fmt.Println(i)
		k = i
	} else {
		//fmt.Println(j)
		k = j
	}
}

func main() {
	//fmt.Println(max(10, 15) + max(100, 15))
	var c int
	fmt.Println("address of c:", &c)
	max(10, 15, c)
	fmt.Println(c)
}
