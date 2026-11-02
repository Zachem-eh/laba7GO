package product

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

func init() {
	first, _ := NewProduct("Пакет для покупок", 12)
	AddProduct(first)
	second, _ := NewProduct("Левый тапок", 49)
	AddProduct(second)
	third, _ := NewProduct("Правый тапок", 51)
	AddProduct(third)
}

type Product struct {
	ID    uint
	Name  string
	Price float64
}

func (pr Product) String() string {
	return fmt.Sprintf("ID товара: %d\nНазвание товара: %s\nЦена: %.2f", pr.ID, pr.Name, pr.Price)
}

var AllProducts = make(map[uint]Product)

func GetNextID() uint {
	var maxID uint = 0
	for id := range AllProducts {
		if id > maxID {
			maxID = id
		}
	}
	return maxID + 1
}

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

func GetProductByID(id uint) (Product, error) {
	p, exists := AllProducts[id]
	if !exists {
		return Product{}, fmt.Errorf("товар с ID = %d не найден", id)
	}
	return p, nil
}

// Добавление
func AddProduct(p *Product) error {
	if _, exists := AllProducts[p.ID]; exists {
		return fmt.Errorf("товар с ID %d уже существует", p.ID)
	}
	AllProducts[p.ID] = *p
	return nil
}

func PrintAllProducts() {
	for _, p := range AllProducts {
		fmt.Println(p)
		fmt.Println("--------")
	}
}

func InputNewProduct(reader *bufio.Reader) error {
	fmt.Println("\nВведите название:")
	name, _ := reader.ReadString('\n')
	name = strings.TrimSpace(name)

	fmt.Println("\nВведите цену:")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	price, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return fmt.Errorf("Некорректный ввод! Пожалуйста, введите для цены положительное число.")
	}
	new_pr, err := NewProduct(name, price)
	if err != nil {
		return err
	}
	AddProduct(new_pr)
	fmt.Println("\nДобавлен товар:")
	fmt.Println(new_pr)
	return nil
}
