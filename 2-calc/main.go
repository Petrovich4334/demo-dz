package main

import (
	"errors"
	"fmt"
)

// Основная функция программы
func main() {
	var op string      // Переменная для хранения операции
	var numsStr string // Переменная для хранения строки с числами

	// Запрос ввода операции
	fmt.Print("Введите операцию (AVG, SUM, MED): ")
	fmt.Scanln(&op)
	op = trimSpace(op) // Удаление пробелов из строки операции
	op = toUpper(op)   // Преобразование в верхний регистр

	// Запрос ввода чисел
	fmt.Print("Введите числа, разделенные запятой: ")
	fmt.Scanln(&numsStr)
	numsStr = trimSpace(numsStr) // Удаление пробелов из строки чисел

	// Парсинг чисел из строки
	numbers := parseNumbers(numsStr)
	if len(numbers) == 0 {
		fmt.Println("Ошибка: Не предоставлено валидных чисел")
		return
	}

	var result float64 // Переменная для хранения результата
	// Выбор операции и вычисление результата
	switch op {
	case "SUM":
		result = sumInts(numbers) // Вычисление суммы
	case "AVG":
		result = avgInts(numbers) // Вычисление среднего
	case "MED":
		result = medInts(numbers) // Вычисление медианы
	default:
		fmt.Println("Ошибка: Неверная операция. Используйте AVG, SUM или MED")
		return
	}

	// Вывод результата
	fmt.Printf("Результат: %.2f\n", result)
}

// Функция для удаления пробелов в начале и конце строки
func trimSpace(s string) string {
	start := 0
	for start < len(s) && s[start] == ' ' {
		start++
	}
	end := len(s) - 1
	for end >= start && s[end] == ' ' {
		end--
	}
	return s[start : end+1]
}

// Функция для преобразования строки в верхний регистр
func toUpper(s string) string {
	result := make([]byte, len(s))
	for i, c := range s {
		if c >= 'a' && c <= 'z' {
			result[i] = byte(c - 32) // Преобразование строчной буквы в заглавную (ASCII)
		} else {
			result[i] = byte(c)
		}
	}
	return string(result)
}

// Функция для парсинга строки чисел, разделенных запятыми
func parseNumbers(s string) []int64 {
	var nums []int64 // Слайс для хранения чисел
	i := 0
	for i < len(s) {
		j := i
		// Поиск следующей запятой
		for j < len(s) && s[j] != ',' {
			j++
		}
		numStr := s[i:j] // Подстрока с числом
		num, err := parseInt(numStr)
		if err == nil {
			nums = append(nums, num) // Добавление валидного числа
		} else if len(numStr) > 0 {
			// Предупреждение о пропуске неверного числа
			fmt.Printf("Предупреждение: Неверное число '%s' пропущено\n", numStr)
		}
		if j < len(s) {
			i = j + 1 // Переход к следующему числу
		} else {
			break
		}
	}
	return nums
}

// Функция для парсинга строки в целое число
func parseInt(s string) (int64, error) {
	if len(s) == 0 {
		return 0, errors.New("пустая строка") // Ошибка для пустой строки
	}
	var num int64 = 0
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c < '0' || c > '9' {
			return 0, errors.New("недопустимый символ") // Ошибка для неверного символа
		}
		num = num*10 + int64(c-'0') // Построение числа
	}
	return num, nil
}

// Функция для вычисления суммы чисел
func sumInts(nums []int64) float64 {
	var s float64 = 0
	for _, n := range nums {
		s += float64(n)
	}
	return s
}

// Функция для вычисления среднего арифметического
func avgInts(nums []int64) float64 {
	if len(nums) == 0 {
		return 0
	}
	return sumInts(nums) / float64(len(nums))
}

// Функция для вычисления медианы
func medInts(nums []int64) float64 {
	if len(nums) == 0 {
		return 0
	}
	sortInts(nums) // Сортировка массива
	n := len(nums)
	if n%2 == 1 {
		return float64(nums[n/2]) // Медиана для нечетного количества
	}
	m1 := float64(nums[n/2-1])
	m2 := float64(nums[n/2])
	return (m1 + m2) / 2 // Медиана для четного количества
}

// Функция сортировки пузырьком для целых чисел
func sortInts(a []int64) {
	n := len(a)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j] // Обмен элементов
			}
		}
	}
}
