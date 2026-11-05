// Команда shop — консольный магазин: ассортимент, корзина, промокоды и чек.
//
// Запуск из корня модуля:
//
//	go run ./cmd/shop
package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"

	"laba7/cart"
	"laba7/discount"
	_ "laba7/discount/promo" // blank import: init() пакета регистрирует промокоды
	catalog "laba7/internal/storage"
	"laba7/product"
	"laba7/receipt"
)

func clearScreen() {
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd", "/c", "cls")
	} else {
		cmd = exec.Command("clear")
	}
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	promoCode := "" // введённый промокод; пустая строка — промокода нет
	for {
		fmt.Println("Для начала работы пожалуйста, нажмите Enter...")
		reader.ReadString('\n')
		clearScreen()

		fmt.Println("Меню команд:")
		fmt.Println("1 - Вывести ассортимент")
		fmt.Println("2 - Добавить товар в ассортимент")
		fmt.Println("3 - Добавить товар в корзину")
		fmt.Println("4 - Вывести корзину")
		fmt.Println("5 - Очистить корзину")
		fmt.Println("6 - Ввести промокод")
		fmt.Println("0 - Завершить работу")
		fmt.Println("Выберите пункт из меню управления: ")

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		a, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Некорректный ввод! Пожалуйста, введите число.")
			continue
		}

		switch a {
		case 0:
			return
		case 1:
			fmt.Println("\nАссортимент продуктов:")
			for _, p := range catalog.All() {
				fmt.Println(p)
				fmt.Println("--------")
			}
		case 2:
			p, err := product.ReadProduct(reader)
			if err != nil {
				fmt.Println(err)
				continue
			}
			p = catalog.Add(p)
			fmt.Println("\nДобавлен товар:")
			fmt.Println(p)
		case 3:
			cart.MakeOrder(reader)
		case 4:
			var modifier discount.PriceModifier // nil, если промокода нет
			if promoCode != "" {
				// Каждый раз берём новый модификатор: у ограниченных
				// скидок есть счётчик, и он должен начинаться с нуля.
				modifier, err = discount.ByCode(promoCode)
				if err != nil {
					fmt.Println(err)
					continue
				}
			}
			receipt.PrintBasket(modifier, promoCode)
		case 5:
			cart.DeleteBasket()
			promoCode = ""
			fmt.Println("Корзина очищена.")
		case 6:
			fmt.Println("Доступные промокоды:", strings.Join(discount.Codes(), ", "))
			fmt.Println("Введите промокод:")
			code, _ := reader.ReadString('\n')
			code = strings.ToUpper(strings.TrimSpace(code))
			if _, err := discount.ByCode(code); err != nil {
				fmt.Println(err)
				continue
			}
			promoCode = code
			fmt.Println("Промокод применён:", promoCode)
		default:
			fmt.Println("Нет такого пункта меню!")
		}
	}
}
