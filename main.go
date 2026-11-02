package main

import (
	"bufio"
	"fmt"
	"laba7/cart"
	"laba7/product"
	"laba7/receipt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
)

func clear_screen() {
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
	for {
		fmt.Println("Для начала работы пожалуйста, нажмите Enter...")
		reader.ReadString('\n')
		clear_screen()

		fmt.Println("Меню команд:")
		fmt.Println("1 - Вывести ассортимент")
		fmt.Println("2 - Добавить товар в ассортимент")
		fmt.Println("3 - Добавить товар в корзину")
		fmt.Println("4 - Вывести корзину")
		fmt.Println("5 - Очистить корзину")
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
			product.PrintAllProducts()
		case 2:
			fmt.Println("\nВведите название:")
			name, _ := reader.ReadString('\n')
			name = strings.TrimSpace(name)

			fmt.Println("\nВведите цену:")
			input, _ := reader.ReadString('\n')
			input = strings.TrimSpace(input)
			price, err := strconv.ParseFloat(input, 64)
			if err != nil {
				fmt.Println("Некорректный ввод! Пожалуйста, введите число.")
				continue
			}
			new_pr, err := product.NewProduct(name, price)
			if err != nil {
				fmt.Println(err)
				continue
			}
			product.AddProduct(new_pr)
			fmt.Println("\nДобавлен товар:")
			fmt.Println(new_pr)
		case 3:
			cart.MakeOrder(reader, product.AllProducts)
		case 4:
			receipt.PrintBasket()
		case 5:
			cart.DeleteBasket()
			fmt.Println("Корзина очищена.")
		default:
			fmt.Println("Нет такого пункта меню!")
		}
	}
}
