package main

import "fmt"

const gCap int = 5 //total array elements
var gLen int       //int number of items in aray currently
//var gGroceries [gCap]string //array of strings

var gGroceries []string //slice of groceries

func addGrocery(a ...string) {

	//fmt.Println("Capacity:", cap(gGroceries))
	for _, d := range a {
		gGroceries = append(gGroceries, d)
	}

	//gGroceries = append(gGroceries, a)

	/*if gLen < gCap {
		gGroceries[gLen] = a
		gLen++
	} else {
		fmt.Println("Error!! Reached array's limit!")
	}*/
}

/*func listGroceries() {
	fmt.Println("Grocery list is as follows:")
	for i := 0; i < gLen; i++ {
		fmt.Println(gGroceries[i])
	}
}*/

func listGroceries() {
	/*for i, d := range gGroceries {
		fmt.Println(i, d)
	}*/

	for _, d := range gGroceries {
		fmt.Println(d)
	}
}

func main() {
	//listGroceries()
	addGrocery("Coffee")
	addGrocery("Milk")
	addGrocery("Pizza")
	addGrocery("Pizza")
	addGrocery("Pizza")
	addGrocery("Pizza")
	addGrocery("Pizza")
	listGroceries()
}
