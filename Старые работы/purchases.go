package main

import (
	"errors"
	"fmt"
)

func analyzePurchases(prices []float64) (sum float64, avg float64, count int, err error) {

	if len(prices) == 0 {
		return 0, 0, 0, errors.New("список покупок пуст")
	}

	for _, price := range prices {
		sum = sum + price
	}
	count = len(prices)
	avg = sum / float64(count)
	return sum, avg, count, nil
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

	sum, avg, count, err := analyzePurchases(prices)
	if err != nil {
		fmt.Println("Ошибка:", err)
	} else {
		fmt.Println("Cумма цен:", sum)
		fmt.Println("Средняя цена:", avg)
		fmt.Println("Кол-во товаров:", count)
	}
}
