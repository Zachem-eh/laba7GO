package main

import (
	"bufio"
	"fmt"
	"laba7/product"
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
	first, _ := product.NewProduct("Пакет для покупок", 12)
	product.AddProduct(first)

	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Println("Для начала работы пожалуйста, нажмите Enter...")
		reader.ReadString('\n')
		clear_screen()

		fmt.Println("Меню команд:")
		fmt.Println("1 - Вывести ассортимент")
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
			fmt.Println("Ассортимент продуктов:")
			product.PrintAllProducts()
		default:
			fmt.Println("Нет такого пункта меню!")
		}
	}
}
