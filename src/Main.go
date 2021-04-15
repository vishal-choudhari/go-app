package main

import (
	"fmt"
)

func main() {
	fractionA := Fraction{
		Numerator:   5,
		Denominator: 25,
	}

	fractionB := Fraction{
		Numerator:   5,
		Denominator: 4,
	}

	fmt.Printf("Fraction A - %s \n", fractionA.PrintFraction())
	fmt.Printf("Fraction B - %s \n", fractionB.PrintFraction())

	fractionC := fractionA.MultiplyByFraction(fractionB)

		fmt.Printf("Result of Multiplication of Fraction A and Fraction B is - %s \n", fractionC.PrintFraction())
}
