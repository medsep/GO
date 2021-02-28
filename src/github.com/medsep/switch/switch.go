package main

import "fmt"

func main() {
	/*for day := 1; day <= 12; day++ {
		if day == 12 {
			fmt.Println("12 drummers drumming")
		} else if day == 11 {
			fmt.Println("11 drummers drumming")
		}
	}*/

	for day := 1; day <= 12; day++ {
		switch day {
		case 12:
			fmt.Println("12th Day")
		case 11:
			fmt.Println("11th Day")
		case 10:
			fmt.Println("10th Day")
		case 9:
			fmt.Println("9th Day")
		case 8:
			fmt.Println("8th Day")
		case 7:
			fmt.Println("7th Day")
		case 6:
			fmt.Println("6th Day")
		case 5:
			fmt.Println("5th Day")
		case 4:
			fmt.Println("4th Day")
		case 3:
			fmt.Println("3 Day")
		case 2:
			fmt.Println("2 Day")
		case 1:
			fmt.Println("1 Day")
		}
	}
}
