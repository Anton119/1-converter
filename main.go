package main

import (
	"errors"
	"fmt"
)

var exchangeRates = map[string]map[string]float64{
	"USD": {"RUB": 80.0, "EUR": 0.88},
	"EUR": {"RUB": 91.0, "USD": 1.13},
	"RUB": {"USD": 84.0, "EUR": 94.6},
}

var currencyOptions = map[int]string{
	1: "USD",
	2: "EUR",
	3: "RUB",
}

func main() {
	fmt.Println("Добро пожаловать в конвертер валют!")
Menu:
	for {
		fmt.Println("Выберите валюту, которую вы хотите конвертировать: 1)USD 2)EUR 3)RUB")
		userChoice, err := getUserCurrency()
		if err != nil {
			fmt.Println(err)
			break Menu
		}

		fmt.Println("Введите сумму для обмена: ")
		sum := getUserSum()

		currency, convertTo_1, convertTo_2, err := availableCurrencyToConvert(&exchangeRates, &currencyOptions, userChoice)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Printf("Выберите валюту для конвертации: 1)%s 2)%s\n", convertTo_1, convertTo_2)
		outcome, err := convertTo(currency, sum, convertTo_1, convertTo_2, &exchangeRates)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("Результат:", outcome)

	}
}

func getUserCurrency() (int, error) {
	var input int
	fmt.Scan(&input)

	if input < 0 || input > 3 {
		return 0, errors.New("Выбор должен быть от 1 до 3")
	}
	return input, nil
}

func getUserSum() int {
	var sum int
	fmt.Scan(&sum)
	return sum
}

// возвращает валюту для конвертации и 2 валюты во что конвертировать
func availableCurrencyToConvert(exchangeRates *map[string]map[string]float64, currencyOptions *map[int]string, userChoice int) (string, string, string, error) {

	convertedCurrency := make([]string, 0, 2)

	currency, ok := (*currencyOptions)[userChoice]
	if !ok {
		return "err", "err", "err", errors.New("Невалидное значение выбора")
	}

	convertRates, ok := (*exchangeRates)[currency]
	if !ok {
		return "err", "err", "err", errors.New("Валюта не найдена")
	}

	for key, _ := range convertRates {
		convertedCurrency = append(convertedCurrency, key)
	}
	return currency, convertedCurrency[0], convertedCurrency[1], nil

}

func convertTo(currency string, sum int, convertTo_1 string, convertTo_2 string, exchangeRates *map[string]map[string]float64) (float64, error) {
	choices := map[int]string{
		1: convertTo_1,
		2: convertTo_2,
	}

	var choice int
	fmt.Scan(&choice)

	if choice < 0 || choice > 2 {
		return 0.0, errors.New("Выбери от 1 до 2")
	}

	covertionCurrecny := choices[choice]
	res := 0.0
	for key, val := range (*exchangeRates)[currency] {
		if key == covertionCurrecny {
			res = float64(sum) * val
		}
	}
	return res, nil
}
