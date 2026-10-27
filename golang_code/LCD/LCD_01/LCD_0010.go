package main

import "fmt"

type LCD_0010_Calculator struct {
	helper LCD_0020_MathHelper
}

func (c LCD_0010_Calculator) CalculateLCD(a, b int) int {
	if a == 0 || b == 0 {
		return 0
	}
	absA := a
	if absA < 0 {
		absA = -absA
	}
	absB := b
	if absB < 0 {
		absB = -absB
	}
	return (absA * absB) / c.helper.gcd(absA, absB)
}

func main() {
	calc := LCD_0010_Calculator{helper: LCD_0020_MathHelper{}}
	fmt.Printf("LCD of 12 and 18 is: %d\n", calc.CalculateLCD(12, 18))
	fmt.Printf("LCD of 5 and 7 is: %d\n", calc.CalculateLCD(5, 7))
}
