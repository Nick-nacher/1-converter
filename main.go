package main

import (
	"fmt"
	"math"
)

func main() {

	res := getConvertVal(getUserEnterValue())
	fmt.Println(math.Round(res))

}

func getConvertVal(val1, val2 string, sum float64) float64 {
	const eur float64 = 0.87
	const usd float64 = 78.44
	var convert float64

	switch val1 {
	case "RUB":
		switch val2 {
		case "USD":
			convert = sum / usd
		case "EUR":
			convert = sum / usd * eur
		case "RUB":
			convert = sum
		}
	case "USD":
		switch val2 {
		case "USD":
			convert = sum
		case "EUR":
			convert = sum * eur
		case "RUB":
			convert = sum * usd
		}
	case "EUR":
		switch val2 {
		case "USD":
			convert = sum * eur
		case "EUR":
			convert = sum
		case "RUB":
			convert = sum * usd * eur

		}
	default:
		convert = 0
	}
	return convert
}

func getUserEnterValue() (string, string, float64) {

	var val1 string
	var val2 string
	var sum float64

BEGIN:
	fmt.Println("Введите код валюты 1 (USD/EUR/RUB):")
	fmt.Scan(&val1)
	cont, hint := getUserHint(val1, val2, 1)
	if cont {
	CONTINUE:
		fmt.Println("Введите код валюты 2", hint)
		fmt.Scan(&val2)

		cont, hint := getUserHint(val1, val2, 2)
		if cont {
		ENTERSUM:
			fmt.Println("Введите конвертируемую сумму:")
			fmt.Scan(&sum)
			if !checkUserSum(sum) {
				fmt.Println("Сумма указана неверно")
				goto ENTERSUM
			}
		} else {
			fmt.Println(hint)
			goto CONTINUE
		}
	} else {
		fmt.Println(hint)
		goto BEGIN
	}

	return val1, val2, sum
}

func checkUserSum(sum float64) bool {
	if sum > 0 {
		return true
	} else {
		return false
	}
}
func getUserHint(val1, val2 string, currency int) (bool, string) {
	var par1 bool
	var par2 string
	switch currency {
	case 1:
		switch val1 {
		case "USD":
			par1 = true
			par2 = "EUR/RUB"
		case "EUR":
			par1 = true
			par2 = "USD/RUB"
		case "RUB":
			par1 = true
			par2 = "USD/EUR"
		default:
			par1 = false
			par2 = "Код валюты введен неверно"
		}
	case 2:
		switch {
		case val1 == "USD" && val2 == "RUB":
			par1 = true
			par2 = "Ввод верный"
		case val1 == "USD" && val2 == "EUR":
			par1 = true
			par2 = "Ввод верный"
		case val1 == "EUR" && val2 == "RUB":
			par1 = true
			par2 = "Ввод верный"
		case val1 == "EUR" && val2 == "USD":
			par1 = true
			par2 = "Ввод верный"
		case val1 == "RUB" && val2 == "USD":
			par1 = true
			par2 = "Ввод верный"
		case val1 == "RUB" && val2 == "EUR":
			par1 = true
			par2 = "Ввод верный"
		default:
			par1 = false
			par2 = "Код валюты введен неверно"
		}

	}
	return par1, par2
}
