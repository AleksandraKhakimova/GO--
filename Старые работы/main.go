package main

import (
	"errors"
	"fmt"
)

func filterByPrice(prices []float64, maxPrice float64) ([]float64, error) {

	if len(prices) == 0 {
		return nil, errors.New("список пуст")
	}
	var result []float64
	for _, price := range prices {
		if price <= maxPrice {
			result = append(result, price)
		}
	}

	if len(result) == 0 {
		return nil, fmt.Errorf("нет товаров дешевле %.2f", maxPrice)
	}
	return result, nil
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
	var maxPrice float64
	fmt.Print("Какой лимит по цене?")
	fmt.Scan(&maxPrice)

	result, err := filterByPrice(prices, maxPrice)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Товары подходящие по цене:", result)
	}
}
