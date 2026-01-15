package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {

	res, err := getConvertVal(getUserEnterValue())
	if err == nil {
		fmt.Println(math.Round(res))
	} else {
		fmt.Println(err)
	}

}

func getConvertVal(val1, val2 string, sum float64) (float64, error) {
	const eur float64 = 0.87
	const usd float64 = 78.44

	var convert float64

	listVal := getRateVal()
	mapVal := listVal[val1]
	rate := mapVal[val2]
	if rate <= 0 {
		return 0, errors.New("тариф для расчета не найден")
	} else {
		convert = sum * rate
		return convert, nil
	}
	//convert = sum * rate

	/*switch val1 {
	case "RUB":
		switch val2 {
		case "USD":
			convert = sum / usd // 4
		case "EUR":
			convert = sum / usd * eur // 5
		case "RUB":
			convert = sum // 1
		}
	case "USD":
		switch val2 {
		case "USD":
			convert = sum // 1
		case "EUR":
			convert = sum * eur // 2
		case "RUB":
			convert = sum * usd // 3
		}
	case "EUR":
		switch val2 {
		case "USD":
			convert = sum * eur // 2
		case "EUR":
			convert = sum // 1
		case "RUB":
			convert = sum * usd * eur // 6

		}
	default:
		convert = 0
	}*/
	//return convert
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

func getRateVal() map[string]map[string]float64 {
	listVal := map[string]map[string]float64{}

	rub := map[string]float64{}
	rub["USD"] = 0.07844
	rub["EUR"] = 0.09016

	usd := map[string]float64{}
	usd["RUB"] = 78.44
	usd["EUR"] = 0.87

	eur := map[string]float64{}
	eur["RUB"] = 90.16
	eur["USD"] = 1.15

	listVal["RUB"] = rub
	listVal["USD"] = usd
	listVal["EUR"] = eur

	return listVal

}
