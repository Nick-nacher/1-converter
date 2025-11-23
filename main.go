package main

import "fmt"

func main() {
	const rub float64 = 79.02
	const eur float64 = 0.87

	const convert float64 = rub / eur
	fmt.Println(convert)

}
