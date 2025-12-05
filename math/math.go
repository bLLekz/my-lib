package math

// Сложение
func Add(a, b int) int {
	return a + b
}

// Умножение
func Multiply(a, b int) int {
	return a * b
}

// Калькулятор
type Calculator struct {
	LastResult int
}

// Сложение всех чисел
func (c *Calculator) Sum(numbers ...int) int {
	total := 0
	for _, n := range numbers {
		total += n
	}
	c.LastResult = total
	return total
}
