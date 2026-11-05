package product

import "fmt"

// Item — товар вместе с количеством.
type Item struct {
	Product Product
	Count   uint
}

// NewItem создаёт Item из товара p и количества count.
// Количество должно быть больше нуля.
func NewItem(p Product, count int) (*Item, error) {
	if count <= 0 {
		return nil, fmt.Errorf("количество должно быть больше нуля")
	}

	return &Item{
		Product: p,
		Count:   uint(count),
	}, nil
}

// String реализует fmt.Stringer.
func (item Item) String() string {
	return fmt.Sprintf("Товар: %s\nКоличество: %d", item.Product.Name, item.Count)
}
