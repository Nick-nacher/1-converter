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
	fmt.Println("Введите код валюты 1 (RUB/EUR):")
	fmt.Scan(&val1)
	fmt.Println("Введите код валюты 2 (RUB/EUR):")
	fmt.Scan(&val2)
	fmt.Println("Введите конвертируемую сумму:")
	fmt.Scan(&sum)
	return val1, val2, sum
}
