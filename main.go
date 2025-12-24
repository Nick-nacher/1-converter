package main

import "fmt"

func main() {

	getUserEnterValue()

}

func getConvertVal(val1, val2 string, sum float64) {

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

// if val1 != nil  {
// 	if val1 != "EUR" || val1 !="USD"{
//     return false
// }else{ if val2 != nil {
// 	if val2 != "EUR" || val2 !="USD"{

// }

// }
// 	return true}
// 	} else if par == 2{}
// 	return true
// }
