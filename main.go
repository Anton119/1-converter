package main

import (
	"errors"
	"fmt"
	"strconv"
)

const EurToRub = 91.2
const EurToUsd = 1.13
const UsdToRub = 80.0
const UsdToEur = 0.88
const RubToUsd = 84.0
const RubToEur = 94.6

const EUR = "EUR"
const RUB = "RUB"
const USD = "USD"

func main() {
	fmt.Println("Добро пожаловать в конвертер валют")
	for {
		fmt.Printf("Выберите название исходной валюты: 1)%s, 2)%s, 3)%s: ", USD, EUR, RUB)
		currency, err := getUserCurrency()
		if err != nil {
			fmt.Println(currency, err)
		}

		fmt.Println("Введите сумму для конвертации: ")
		sum, err := getSum()
		if err != nil {
			fmt.Println(sum, err)
		}

		covertTo_1, convertTo_2, err := availableCurrencyToConver(currency)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Выберите валюту для конвертации: 1)%s, 2)%s: ", covertTo_1, convertTo_2)
		convertCurrencyChoice, err := getCurrecnyToConvert()
		if err != nil {
			fmt.Println("", err)
		}

		currencyToConvert := ""
		switch convertCurrencyChoice {
		case "1":
			currencyToConvert = covertTo_1
		case "2":
			currencyToConvert = convertTo_2
		}

		res := countExchange(currency, sum, currencyToConvert)

		fmt.Printf("Результат вашей операции: %.2f\n", res)

	}
}

func getUserCurrency() (string, error) {
	var input string
	fmt.Scan(&input)
	var currency string

	choice, _ := strconv.Atoi(input)
	switch choice {
	case 1:
		currency = USD
	case 2:
		currency = EUR
	case 3:
		currency = RUB
	default:
		return "", errors.New("INVALID_FORMAT")
	}
	return currency, nil
}

func getSum() (int, error) {
	var sum int
	fmt.Scan(&sum)
	return sum, nil
}

func availableCurrencyToConver(currency string) (string, string, error) {

	currencies := []string{EUR, RUB, USD}
	convertedCurrencies := []string{}
	for i := 0; i < len(currencies); i++ {
		if currency != currencies[i] {
			convertedCurrencies = append(convertedCurrencies, currencies[i])
		}
	}

	return convertedCurrencies[0], convertedCurrencies[1], nil

}

func getCurrecnyToConvert() (string, error) {
	var input string
	fmt.Scan(&input)
	return input, nil

}

func countExchange(currency string, sum int, currencyToConv string) float64 {
	res := float64(sum)
	switch currency {
	case USD:
		switch currencyToConv {
		case RUB:
			return res * UsdToRub
		case EUR:
			return res * UsdToEur
		}
	case EUR:
		switch currencyToConv {
		case USD:
			return res * EurToUsd
		case RUB:
			return res * EurToRub
		}
	case RUB:
		switch currencyToConv {
		case USD:
			return res * RubToUsd
		case EUR:
			return res * RubToEur
		}
	}
	return 0.0
}

// тестирую коммиты
