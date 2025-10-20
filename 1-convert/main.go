package main

import (
	"errors"
	"fmt"
	"strings"
)

var exchangeRates = map[string]float64{
	"USD": 1.0,   // 1 USD = 1 USD (базовая валюта)
	"EUR": 0.85,  // 1 USD = 0.85 EUR
	"RUB": 73.50, // 1 USD = 73.50 RUB
}

func main() {
	fmt.Println("__Калькулятор валют__")

	fromCurrency, err := getCurrencyInput("Введите исходную валюту (USD, EUR, RUB): ")
	if err != nil {
		fmt.Println("Ошибка ввода валюты:", err)
		return
	}

	amount, err := getAmountInput("Введите сумму:")
	if err != nil {
		fmt.Println("Ошибка ввода суммы:", err)
		return
	}

	toCurrency, err := getCurrencyInput("Введите целевую валюту (USD, EUR, RUB): ")
	if err != nil {
		fmt.Println("Ошибка ввода валюты:", err)
		return
	}

	//Проверка одинаковых валют: чтобы избежать ненужных вычислений и сразу вернуть исходную сумму
	if fromCurrency == toCurrency {
		fmt.Printf("%.2f %s = %.2f %s (валюты одинаковы)\n", amount, fromCurrency, amount, toCurrency)
		return
	}

	result, err := convertCurrency(amount, fromCurrency, toCurrency)
	if err != nil {
		fmt.Println("Ошибка конвертации:", err)
		return
	}

	fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)
}

// getCurrencyInput запрашивает валюту и проверяет её допустимость
func getCurrencyInput(prompt string) (string, error) {
	var currency string

	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&currency)
		if err != nil {
			return "", errors.New("ошибка чтения ввода")
		}

		currency = strings.ToUpper(currency)

		if _, exists := exchangeRates[currency]; exists {
			return currency, nil
		}
		fmt.Println("Недопустимая валюта, используйте USD, EUR или RUB")
	}
}

// getAmountInput запрашивает сумму и проверяет её корректность
func getAmountInput(prompt string) (float64, error) {
	var amount float64

	for {
		fmt.Print(prompt)

		_, err := fmt.Scan(&amount)
		if err != nil {
			fmt.Println("Ошибка ввода, попробуйте снова")
			continue
		}

		if amount < 0 {
			fmt.Println("Сумма не может быть отрицательной, попробуйте снова")
			continue
		}

		return amount, nil
	}
}

func convertCurrency(amount float64, fromCurrency string, toCurrency string) (float64, error) {

	fromRate, fromExists := exchangeRates[fromCurrency]

	if !fromExists {
		return 0, fmt.Errorf("исходная валюта %s не поддерживается", fromCurrency)
	}

	toRate, toExists := exchangeRates[toCurrency]
	// Если целевая валюта не найдена, возвращаем ошибку
	if !toExists {
		return 0, fmt.Errorf("целевая валюта %s не поддерживается", toCurrency)
	}

	result := (amount / fromRate) * toRate

	return result, nil
}
