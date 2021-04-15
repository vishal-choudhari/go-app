package main

import (
	"fmt"
)

type Fraction struct {
	Numerator, Denominator int
}

func (f Fraction) PrintFraction() string {
	return fmt.Sprintf("%d/%d", f.Numerator, f.Denominator)
}

// Reduce makes given fraction reduced.
func (f Fraction) Reduce() <-chan Fraction {
	ch := make(chan Fraction)

	go func() {
		defer close(ch)
		if f.Numerator != 0 && f.Denominator !=0 {
			gcd := GCDEuclidean(f.Numerator, f.Denominator)
			f.Numerator /= gcd
			f.Denominator /= gcd
		} else {
			f.Numerator = 0
			f.Denominator = 0
		}
		ch <- f
	}()

	return ch
}


// multiplies fraction by given fraction by first reducing both the fractions and also reducing the result.
func (f Fraction) MultiplyByFraction(m Fraction) Fraction {

	if m.Denominator != 0 && f.Denominator != 0 {
		fCh, mCh := f.Reduce(), m.Reduce()
		f, m = <-fCh, <-mCh

		f.Numerator *= m.Numerator
		f.Denominator *= m.Denominator

		f = <-f.Reduce()

		return f

	} else {
		fmt.Printf("Denominator is zero for one of the fraction, so skipping multiplication. \n ")
		f.Denominator = 0
		f.Numerator = 0
		return f
	}

}

// calculates GCD by Euclidian algorithm.
func GCDEuclidean(a, b int) int {

	if a == 0 {
		return b
	} else if b == 0 {
		return a
	} else {
		for a != b {
			if a > b {
				a -= b
			} else {
				b -= a
			}
		}
		return a
	}
}
