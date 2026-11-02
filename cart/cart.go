package cart

import (
	"bufio"
	"fmt"
	"io"
	"laba7/product"
	"strconv"
	"strings"
)

var basket []product.Product

func MakeOrder(reader *bufio.Reader, products map[uint]product.Product) {
	fmt.Println("Весь ассортимент продуктов:")
	product.PrintAllProducts()
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
		item, err := product.GetProductByID(uint(id))
		if err != nil {
			fmt.Println(err)
			continue
		}
		basket = append(basket, item)
	}
}
func GetBasket() []product.Product {
	return basket
}
func DeleteBasket() {
	basket = nil
}
