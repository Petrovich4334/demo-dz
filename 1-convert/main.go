package main

import "fmt"

func main() {
	fmt.Println("__Калькулятор валют__")

	const usdToEur = 0.85
	const usdToRub = 73.50

	var amount float64
	fmt.Print("Введите сумму в USD: ")
	fmt.Scan(&amount)

	eurAmount := amount * usdToEur
	rubAmount := amount * usdToRub

	fmt.Printf(`Результаты конвертации:
	%.2f USD = %.2f EUR
	%.2f USD = %.2f RUB
	`, amount, eurAmount, amount, rubAmount)
}
