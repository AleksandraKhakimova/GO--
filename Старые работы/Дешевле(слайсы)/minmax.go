package main

import (
	"errors"
	"fmt"
)

func findMinMax(prices []float64) (min float64, max float64, err error) {

	if len(prices) == 0 {
		return 0, 0, errors.New("список пуст")
	}
	max = prices[0]
	min = prices[0]
	for _, price := range prices {
		if price < min {
			min = price
		}
		if price > max {
			max = price
		}
	}
	return min, max, nil
}

func main() {
	var count int
	fmt.Print("Сколько товаров?")
	fmt.Scan(&count)
	var prices []float64
	for i := 1; i <= count; i++ {
		var price float64
		fmt.Printf("Цена товара %d:", i)
		fmt.Scan(&price)
		prices = append(prices, price)
	}

	min, max, err := findMinMax(prices)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Товары по самой низкой цене:", min)
		fmt.Println("Товары по самой высокой цене:", max)
	}
}
