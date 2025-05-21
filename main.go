package main

import "fmt"

func main() {
	const UsdToEur = 1.13
	const UsdToRub = 80

	eurToRub := UsdToEur * UsdToRub
	fmt.Println(eurToRub)
}
