package main

import "fmt"

func main() {
	fmt.Println("__Калькулятор валют__")

	const usdToEur = 0.85
	const usdToRub = 73.50

	amount := getAmountInput("Введите сумму в USD: ")

	eurAmount := amount * usdToEur
	rubAmount := amount * usdToRub

	fmt.Printf(`Результаты конвертации:
	%.2f USD = %.2f EUR
	%.2f USD = %.2f RUB
	`, amount, eurAmount, amount, rubAmount)
}

func getAmountInput(prompt string) float64 {
	var amount float64
	fmt.Print(prompt)
	fmt.Scan(&amount)
	return amount
}

func convertCurrency(amount float64, fromCurrency string, toCurrency string) float64 {
	return 0
}
