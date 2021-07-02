package main

import (
	"fmt"
)

func main() {
	var n, p, q int
	var dayRemaining int

	fmt.Scanln(&n, &p, &q)

	tempCalculation1 := n % p
	tempCalculation2 := n % q

	dayRemaining = 0
 
	if (tempCalculation1 == 0 && tempCalculation2 != 0) || (tempCalculation2 == 0 && tempCalculation1 != 0) {
		fmt.Println("Kita bertemu hari ini")
	} else if (p != q) && (p - tempCalculation1 > 0 && q - tempCalculation2 > 0) {
		dayRemaining = 0

		tempDayRemainingP := p - tempCalculation1
		tempDayRemainingQ := q - tempCalculation2

		if tempDayRemainingP < tempDayRemainingQ {
			dayRemaining = tempDayRemainingP
		} else {
			dayRemaining = tempDayRemainingQ
		}

		fmt.Println("Kita bertemu", dayRemaining, "hari lagi")
	} else if tempCalculation1 == 0 && tempCalculation2 == 0 {
		fmt.Println("Kita bertemu 7 hari lagi")
	} else if p == q {
		fmt.Println("Kita tidak pernah bertemu")
	}
}