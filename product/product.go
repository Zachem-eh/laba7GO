package product

import (
	"fmt"
)

type Product struct {
	ID    uint
	Name  string
	Price float64
}

var AllProduts = make(map[uint]Product)

func NewProduct(name string, price float64) (*Product, error) {
	if name == "" {
		return nil, fmt.Errorf("Имя товара не должно быть пустым")
	}
	if price < 0 {
		return nil, fmt.Errorf("Цена не может быть отрицательной")
	}

	// Автоматически генерируем ID
	newID := GetNextID()

	return &Product{
		ID:    newID,
		Name:  name,
		Price: price,
	}, nil
}

func (pr Product) String() string {
	return fmt.Sprintf("ID товара: %d\nНазвание товара: %s\nЦена: %.2f", pr.ID, pr.Name, pr.Price)
}

func GetNextID() uint {
	var maxID uint = 0
	for id := range AllProduts {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}

func GetProductByID(id uint) (Product, error) {
	p, exists := AllProduts[id]
	if !exists {
		return Product{}, fmt.Errorf("товар не найден")
	}
	return p, nil
}

// Добавление
func AddProduct(p *Product) error {
	if _, exists := AllProduts[p.ID]; exists {
		return fmt.Errorf("товар с ID %d уже существует", p.ID)
	}
	AllProduts[p.ID] = *p
	return nil
}
