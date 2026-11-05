// Package cart описывает корзину покупателя: добавление товаров из
// ассортимента, очистку и расчёт итоговой стоимости с учётом скидок.
package cart

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"

	"laba7/discount"
	"laba7/internal/storage"
	"laba7/product"
)

var basket []product.Product

// MakeOrder показывает ассортимент, читает из reader ID товаров через пробел
// и добавляет найденные товары в корзину. Если товара с каким-то ID нет,
// печатает ошибку и переходит к следующему ID.
func MakeOrder(reader *bufio.Reader) {
	fmt.Println("Весь ассортимент продуктов:")
	for _, p := range storage.All() {
		fmt.Println(p)
		fmt.Println("--------")
	}
	fmt.Println("\nВведите ID товаров через пробел:")
	// parseIDs разбирает строку вида "1 3 3" в срез ID.
	parseIDs := func(s string) ([]int, error) {
		fields := strings.Fields(s)
		if len(fields) == 0 {
			return nil, fmt.Errorf("не введено ни одного ID")
		}
		ids := make([]int, 0, len(fields))
		for _, field := range fields {
			id, err := strconv.Atoi(field)
			if err != nil {
				return nil, fmt.Errorf("%q не является числом", field)
			}
			ids = append(ids, id)
		}
		return ids, nil
	}

	line, readErr := reader.ReadString('\n')
	if readErr != nil && readErr != io.EOF {
		fmt.Printf("ошибка чтения: %v\n", readErr)
		return
	}
	ids, err := parseIDs(line)
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, id := range ids {
		if id <= 0 {
			fmt.Printf("ID должен быть положительным, получено %d\n", id)
			continue
		}
		item, err := storage.Find(uint(id))
		if err != nil {
			fmt.Println(err)
			continue
		}
		basket = append(basket, item)
	}
}

// GetBasket возвращает копию содержимого корзины.
func GetBasket() []product.Product {
	out := make([]product.Product, len(basket))
	copy(out, basket)
	return out
}

// DeleteBasket очищает корзину.
func DeleteBasket() {
	basket = nil
}

// Total возвращает итоговую стоимость корзины: modifier применяется к цене
// каждого товара по очереди, в порядке добавления в корзину, и полученные
// цены складываются. Если modifier равен nil, цены берутся без изменений.
func Total(modifier discount.PriceModifier) float64 {
	sum := 0.0
	for _, p := range basket {
		price := p.Price
		if modifier != nil {
			price = modifier(price)
		}
		sum += price
	}
	return sum
}
