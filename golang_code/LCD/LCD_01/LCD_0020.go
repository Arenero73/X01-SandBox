package main

type LCD_0020_MathHelper struct{}

func (m LCD_0020_MathHelper) gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}
