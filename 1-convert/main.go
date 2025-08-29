package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("__Калькулятор валют__")

	// Получение всей информации от пользователя
	fromCurrency, err := getCurrencyInput("Введите исходную валюту  (USD, EUR, RUB): ")
	if err != nil {
		fmt.Println("Ошибка ввода валюты:", err)
		return
	}

	amount, err := getAmountInput("Введите сумму: ")
	if err != nil {
		fmt.Println("Ошибка ввода суммы", err)
		return
	}

	toCurrency, err := getCurrencyInput("Введите целевую валюту (USD, EUR, RUB): ")
	if err != nil {
		fmt.Println("Ошибка ввода валюты:", err)
		return
	}

	// Конвертация валюты
	result := convertCurrency(amount, fromCurrency, toCurrency)

	// Вывод результатов конвертации
	fmt.Printf("%.2f %s = %.2f %s\n", amount, fromCurrency, result, toCurrency)
}

// Функция для получения ввода валюты от пользователя
func getCurrencyInput(prompt string) (string, error) {
	var currency string
	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&currency)
		if err != nil {
			return "", errors.New("ошибка чтения ввода")
		}

		// Проверка всех возможных вариантов написания USD
		switch currency {
		case "USD", "usd":
			return "USD", nil
		case "EUR", "eur":
			return "EUR", nil
		case "RUB", "rub":
			return "RUB", nil
		default:
			fmt.Println("Недопустимая валюта, используйте USD, EUR или RUB")
		}
	}
}

// Функция для получения суммы от пользователя
func getAmountInput(prompt string) (float64, error) {
	var amount float64

	for {
		fmt.Print(prompt)
		_, err := fmt.Scan(&amount)

		if err != nil {
			fmt.Println("Ошибка ввода, попробуйте снова.")
			continue
		}

		if amount < 0 {
			fmt.Println("Сумма не может быть отрицательной, поробуйте снова")
			continue
		}
		return amount, nil
	}
}

// Функция для конвертации валют
func convertCurrency(amount float64, fromCurrency string, toCurrency string) float64 {
	const usdToEur = 0.85
	const usdToRub = 73.50

	var inUSD float64

	// Конвертируем исходную валюту в USD используя tagged switch
	switch fromCurrency {
	case "USD":
		inUSD = amount
	case "EUR":
		inUSD = amount / usdToEur
	case "RUB":
		inUSD = amount / usdToRub
	}

	// Конвертируем из USD в целевую валюту используя tagged switch
	switch toCurrency {
	case "USD":
		return inUSD
	case "EUR":
		return inUSD * usdToEur
	case "RUB":
		return inUSD * usdToRub
	}
	return 0
}
