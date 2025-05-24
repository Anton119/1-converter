package main

import (
	"errors"
	"fmt"
	"strings"
)

var exchangeRates = map[string]map[string]float64{
	"USD": {"RUB": 80.0, "EUR": 0.88},
	"EUR": {"RUB": 91.0, "USD": 1.13},
	"RUB": {"USD": 0.012, "EUR": 0.011},
}

var currencyOptions = map[int]string{
	1: "USD",
	2: "EUR",
	3: "RUB",
}

func main() {
	fmt.Println("Добро пожаловать в конвертер валют!")

	for {
		fmt.Println(strings.Repeat("-", 40))
		fmt.Println("Выберите валюту, которую хотите конвертировать:")
		for k, v := range currencyOptions {
			fmt.Printf("%d) %s\n", k, v)
		}

		userChoice, err := getUserCurrency()
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Print("Введите сумму для обмена: ")
		sum, err := getUserSum()
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		currency, targets, err := getTargetCurrencies(userChoice)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Println("Выберите валюту для конвертации:")
		for i, cur := range targets {
			fmt.Printf("%d) %s\n", i+1, cur)
		}

		result, err := convertCurrency(currency, sum, targets)
		if err != nil {
			fmt.Println("Ошибка:", err)
			continue
		}

		fmt.Printf("Результат конвертации: %.2f\n", result)
		break
	}
}

func getUserCurrency() (int, error) {
	var input int
	_, err := fmt.Scanln(&input)
	if err != nil {
		return 0, errors.New("некорректный ввод, введите число")
	}
	if input < 1 || input > 3 {
		return 0, errors.New("выбор должен быть от 1 до 3")
	}
	return input, nil
}

func getUserSum() (int, error) {
	var sum int
	_, err := fmt.Scanln(&sum)
	if err != nil {
		return 0, errors.New("некорректный ввод суммы")
	}
	if sum <= 0 {
		return 0, errors.New("сумма должна быть положительным числом")
	}
	return sum, nil
}

func getTargetCurrencies(userChoice int) (string, []string, error) {
	currency, ok := currencyOptions[userChoice]
	if !ok {
		return "", nil, errors.New("валюта не найдена")
	}

	convertRates, ok := exchangeRates[currency]
	if !ok {
		return "", nil, errors.New("обменные курсы для выбранной валюты отсутствуют")
	}

	targets := make([]string, 0, len(convertRates))
	for cur := range convertRates {
		targets = append(targets, cur)
	}

	if len(targets) == 0 {
		return "", nil, errors.New("нет валют для конвертации")
	}

	return currency, targets, nil
}

func convertCurrency(from string, amount int, targets []string) (float64, error) {
	var choice int
	fmt.Print("Ваш выбор: ")
	_, err := fmt.Scanln(&choice)
	if err != nil {
		return 0, errors.New("некорректный ввод")
	}
	if choice < 1 || choice > len(targets) {
		return 0, errors.New("выбор вне диапазона")
	}

	targetCurrency := targets[choice-1]
	rate := exchangeRates[from][targetCurrency]
	result := float64(amount) * rate
	return result, nil
}
