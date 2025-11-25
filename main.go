// Author: Natnael Debesay
// Version: 1.0.0
// Date 2025-11-25
// Fileoverview: This program calculates the interest value at the end of ten years.

package main
import (
	"fmt"
)

const Years = 10

func calcInterest(prevValue float64, rate float64) float64 {
	rawInterest := prevValue * rate
	roundedValue := rawInterest*100 + 0.5
	return float64(int(roundedValue)) / 100 
}

func calcNewValue(prevValue float64, interest float64) float64 {
	newValue := prevValue + interest
	return float64(int(newValue*100 + 0.5)) / 100
}

func main() {
	var initial float64
	var rate float64

	fmt.Print("Enter Initial value: ")
	fmt.Scan(&initial)

	fmt.Print("Enter Interest rate: ")
	fmt.Scan(&rate)

	// Year 1
	interest1 := calcInterest(initial, rate)
	value1 := calcNewValue(initial, rate)

	// Year 2 
	interest2 := calcInterest(value1, rate)
	value2 := calcNewValue(value1, rate)

	// Year 3
	interest3 := calcInterest(value2, rate)
	value3 := calcNewValue(value2, rate)

	// Year 4
	interest4 := calcInterest(value3, rate)
	value4 := calcNewValue(value3, rate)

	// Year 5
	interest5 := calcInterest(value4, rate)
	value5 := calcNewValue(value4, rate)

	// Year 6
	interest6 := calcInterest(value5, rate)
	value6 := calcNewValue(value5, rate)	

	// Year 7
	interest7 := calcInterest(value6, rate)
	value7 := calcNewValue(value6, rate)

	// Year 8
	interest8 := calcInterest(value7, rate)
	value8 := calcNewValue(value7, rate)

	// Year 9
	interest9 := calcInterest(value8, rate)
	value9 := calcNewValue(value8, rate)

	// Year 10
	interest10 := calcInterest(value9, rate)
	value10 := calcNewValue(value9, rate)

	fmt.Println("\nYear | Initial Value | Interest Gained | Current Value")
	fmt.Println("--------------------------------------------------------")
	fmt.Printf("1 |	$%.2f    |	$%.2f    | $%.2f\n", initial, interest1, value1)
	fmt.Printf("2 |	$%.2f    |	$%.2f    | $%.2f\n", initial, interest2, value2)
	fmt.Printf("3 |	$%.2f    |	$%.2f    | $%.2f\n", initial, interest3, value3)
	fmt.Printf("4 |	$%.2f    |	$%.2f    | $%.2f\n", initial, interest4, value4)
	fmt.Printf("5 |	$%.2f    |	$%.2f    | $%.2f\n", initial, interest5, value5)
	fmt.Printf("6 |	$%.2f    |	$%.2f    | $%.2f\n", initial, interest6, value6)
	fmt.Printf("7 |	$%.2f    |	$%.2f    | $%.2f\n", initial, interest7, value7)
	fmt.Printf("8 |	$%.2f    |	$%.2f    | $%.2f\n", initial, interest8, value8)
	fmt.Printf("9 |	$%.2f    |	$%.2f    | $%.2f\n", initial, interest9, value9)
	fmt.Printf("10 |	$%.2f    |	$%.2f    | $%.2f\n", initial, interest10, value10)

	fmt.Println("\nDone.")
}