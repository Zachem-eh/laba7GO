package product

import "fmt"

type Item struct {
	Product Product
	Count   uint
}

func NewItem(p Product, count int) (*Item, error) {
	if count <= 0 {
		return nil, fmt.Errorf("Количество должно быть больше нуля")
	}

	return &Item{
		Product: p,
		Count:   uint(count),
	}, nil
}

func (item Item) String() string {
	return fmt.Sprintf("Товар: %s\nКоличество: %d", item.Product.Name, item.Count)
}
