// Package storage хранит ассортимент магазина.
//
// Пакет лежит в папке internal, поэтому импортировать его могут только
// пакеты модуля laba7. Сам каталог хранится в неэкспортируемой переменной,
// и работать с ним можно только через функции Add, Find и All.
// Стартовые товары загружаются функцией init при запуске программы.
package storage

import (
	"fmt"
	"sort"

	"laba7/product"
)

// catalog — все товары магазина: ID товара → товар.
// Имя со строчной буквы, поэтому снаружи пакета переменную не видно.
var catalog = make(map[uint]product.Product)

// nextID — ID, который получит следующий добавленный товар.
var nextID uint = 1

// init загружает стартовый ассортимент. Go вызывает её сам один раз до main.
func init() {
	start := []struct {
		name  string
		price float64
	}{
		{"Пакет для покупок", 12},
		{"Левый тапок", 49},
		{"Правый тапок", 51},
	}
	for _, s := range start {
		p, err := product.NewProduct(s.name, s.price)
		if err != nil {
			// Данные зашиты в код: если они неверные, это ошибка программиста.
			panic(err)
		}
		Add(p)
	}
}

// Add присваивает товару p следующий свободный ID, сохраняет его в каталоге
// и возвращает сохранённый товар.
func Add(p product.Product) product.Product {
	p.ID = nextID
	nextID++
	catalog[p.ID] = p
	return p
}

// Find возвращает товар с указанным ID или ошибку, если такого товара нет.
func Find(id uint) (product.Product, error) {
	p, ok := catalog[id]
	if !ok {
		return product.Product{}, fmt.Errorf("товар с ID = %d не найден", id)
	}
	return p, nil
}

// All возвращает все товары каталога, отсортированные по ID.
// Возвращается новый срез, поэтому изменения в нём не затрагивают каталог.
func All() []product.Product {
	list := make([]product.Product, 0, len(catalog))
	for _, p := range catalog {
		list = append(list, p)
	}
	sort.Slice(list, func(i, j int) bool {
		return list[i].ID < list[j].ID
	})
	return list
}
