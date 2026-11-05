// Package product описывает товары магазина: тип Product (товар
// ассортимента) и тип Item (товар с количеством).
//
// Хранением товаров пакет не занимается: каталог лежит в internal/storage.
package product

import (
	"bufio"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

// Product — товар ассортимента.
type Product struct {
	ID    uint
	Name  string
	Price float64
}

// String реализует fmt.Stringer: описание товара в несколько строк.
func (pr Product) String() string {
	return fmt.Sprintf("ID товара: %d\nНазвание товара: %s\nЦена: %.2f", pr.ID, pr.Name, pr.Price)
}

// NewProduct создаёт товар с названием name и ценой price и проверяет данные.
// ID у нового товара равен 0: настоящий ID выдаёт хранилище при добавлении.
func NewProduct(name string, price float64) (Product, error) {
	if name == "" {
		return Product{}, errors.New("имя товара не должно быть пустым")
	}
	if price < 0 {
		return Product{}, errors.New("цена не может быть отрицательной")
	}
	return Product{Name: name, Price: price}, nil
}

// ReadProduct спрашивает у пользователя название и цену товара, читает их
// из reader и возвращает новый товар (ещё без ID).
func ReadProduct(reader *bufio.Reader) (Product, error) {
	fmt.Println("\nВведите название:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("\nВведите цену:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	price, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return Product{}, errors.New("некорректный ввод: цена должна быть числом")
	}
	return NewProduct(name, price)
}
