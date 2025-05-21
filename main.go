package main

import "fmt"

const UsdToEur = 1.13
const UsdToRub = 80

func main() {

	getUserInput()
	eurToRub := UsdToEur * UsdToRub
	fmt.Println(eurToRub)
}

func getUserInput() {
	var input string
	fmt.Scan(&input)
}

func converter(rub float64, eur float64, usd float64) float64 {
	return 0.0
}
