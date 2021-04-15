package main

import (
	"fmt"
	"testing"
)


func TestValidFractionMultipication(t *testing.T) {

	fmt.Println("\n TEST CASE for VALID FRACTION MULTIPLICATIONS.")

	validCases := []struct {
		expected Fraction
		fractionOne Fraction
		fractionTwo Fraction 
	}{
		{Fraction{1, 4}, Fraction{5, 25}, Fraction{5, 4}},
		{Fraction{1, 1}, Fraction{21, 3}, Fraction{9, 63}},
		{Fraction{5, 14}, Fraction{95, 19}, Fraction{1, 14}},
		{Fraction{11, 13}, Fraction{33, 39}, Fraction{1, 1}},
		{Fraction{7, 19}, Fraction{21, 3}, Fraction{1, 19}},
		{Fraction{0, 0}, Fraction{21, 0}, Fraction{1, 19}},
		{Fraction{0, 0}, Fraction{0, 3}, Fraction{1, 19}},
		{Fraction{0, 0}, Fraction{0, 0}, Fraction{1, 19}},
	}

	for _, c  := range validCases {
		// Act
		fmt.Printf(" FractionA : %d/%d , FractionB : %d/%d , ExpectedResult : %d/%d \n ",c.fractionOne.Numerator,c.fractionOne.Denominator,c.fractionTwo.Numerator,c.fractionTwo.Denominator,c.expected.Numerator, c.expected.Denominator)
		actual := c.fractionOne.MultiplyByFraction(c.fractionTwo)
		
		fmt.Printf("Actual Multiplication Result - %d/%d \n ",actual.Numerator, actual.Denominator)

		// Assert
		if actual != c.expected {
			fmt := "Multiplication failed, Expected : %d/%d but got %d/%d" 
			t.Errorf(fmt,c.expected.Numerator, c.expected.Denominator,actual.Numerator, actual.Denominator)
		} else {
			fmt.Print("VALID DATA SET. \n \n")
		}
	}
}

func TestInValidFractionMultipication(t *testing.T) {

	fmt.Println("\n TEST CASE for IN-VALID FRACTION MULTIPLICATIONS.")

	inValidCases := []struct {
		expected Fraction
		fractionOne Fraction
		fractionTwo Fraction 
	}{
		{Fraction{1, 1}, Fraction{5, 25}, Fraction{5, 4}},
		{Fraction{1, 5}, Fraction{21, 3}, Fraction{9, 63}},
		{Fraction{5, 4}, Fraction{95, 19}, Fraction{1, 14}},
		{Fraction{11, 3}, Fraction{33, 39}, Fraction{1, 1}},
		{Fraction{7, 9}, Fraction{21, 3}, Fraction{1, 19}},
	}

	for _, c  := range inValidCases {
		// Act
		fmt.Printf(" FractionA : %d/%d , FractionB : %d/%d , ExpectedResult : %d/%d \n ",c.fractionOne.Numerator,c.fractionOne.Denominator,c.fractionTwo.Numerator,c.fractionTwo.Denominator,c.expected.Numerator, c.expected.Denominator)
		actual := c.fractionOne.MultiplyByFraction(c.fractionTwo)
		
		fmt.Printf("Actual Multiplication Result - %d/%d \n ",actual.Numerator, actual.Denominator)

		// Assert
		if actual == c.expected {
			fmt := "Multiplication failed, Expected : %d/%d but got %d/%d" 
			t.Errorf(fmt,c.expected.Numerator, c.expected.Denominator,actual.Numerator, actual.Denominator)
		} else {
			fmt.Print("INVALID DATA SET. \n \n")
		}
	}
}

func TestGCDEuclideanLogic(t *testing.T) {

	fmt.Println("\n TEST CASE for GCD Euclidean")

	validCases := []struct {
		expected int
		numberOne int
		numberTwo int 
	}{
		{21, 105, 63},
		{11, 33, 11},
		{34, 170, 68},
		{1, 17, 19},
		{105, 210, 315},
		{5, 0, 5},
		{7, 7, 0},
		{0, 0, 0},
	}

	for _, c  := range validCases {
		// Act
		fmt.Printf(" NumberOne : %d , numberTwo : %d , ExpectedResult : %d \n ",c.numberOne,c.numberTwo,c.expected)
		actual := GCDEuclidean(c.numberOne, c.numberTwo)
		
		fmt.Printf("Actual GCD Result - %d \n ",actual)

		// Assert
		if actual != c.expected {
			fmt := "GCD failed, Expected : %d but got %d" 
			t.Errorf(fmt,c.expected,actual)
		} else {
			fmt.Print("VALID DATA SET. \n \n")
		}
	}
}

func TestReduceLogic(t *testing.T) {

	fmt.Println("\n TEST CASE for Reduce fraction")

	validCases := []struct {
		expected Fraction
		given Fraction
	}{
		{Fraction{5, 3}, Fraction{105, 63}},
		{Fraction{3, 1}, Fraction{33, 11}},
		{Fraction{5, 2}, Fraction{1700, 680}},
		{Fraction{17, 19}, Fraction{17, 19}},
		{Fraction{2, 3}, Fraction{210, 315}},
		{Fraction{0, 0}, Fraction{0, 5}},
		{Fraction{0, 0}, Fraction{7, 0}},
	}

	for _, c  := range validCases {
		// Act
		fmt.Printf(" Given : %d/%d , ExpectedResult : %d/%d \n ",c.given.Numerator,c.given.Denominator,c.expected.Numerator, c.expected.Denominator)
		actual := <- c.given.Reduce()
		
		fmt.Printf("Actual Reduce Result - %d/%d \n ",actual.Numerator, actual.Denominator)

		// Assert
		if actual != c.expected {
			fmt := "Reduce failed, Expected : %d/%d but got %d/%d" 
			t.Errorf(fmt,c.expected.Numerator, c.expected.Denominator,actual.Numerator, actual.Denominator)
		} else {
			fmt.Print("VALID DATA SET. \n \n")
		}
	}
}