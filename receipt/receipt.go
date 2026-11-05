// Package receipt печатает чек по содержимому корзины в виде таблицы
// (пакет text/tabwriter).
package receipt

import (
	"fmt"
	"os"
	"text/tabwriter"

	"laba7/cart"
	"laba7/discount"
)

// PrintBasket печатает чек: таблицу товаров корзины и итоговую сумму.
// Если modifier не равен nil (введён промокод code), дополнительно
// печатаются скидка и сумма к оплате с учётом modifier.
func PrintBasket(modifier discount.PriceModifier, code string) {
	fmt.Println("=== Чек ===")
	myBasket := cart.GetBasket()
	if len(myBasket) == 0 {
		fmt.Println("Корзина пуста.")
		return
	}
	// Параметры: os.Stdout (куда писать), 0 (minwidth), 0 (tabwidth), 3 (padding), ' ' (padchar), 0 (flags)
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
	fmt.Fprintln(w, "ID\tНАЗВАНИЕ\tЦЕНА\t")
	fmt.Fprintln(w, "--\t--------\t-----\t")
	for _, p := range myBasket {
		fmt.Fprintf(w, "%d\t%s\t%.2f\t\n", p.ID, p.Name, p.Price)
	}
	w.Flush()

	full := cart.Total(nil)
	fmt.Printf("Сумма: %.2f\n", full)
	if modifier != nil {
		total := cart.Total(modifier)
		fmt.Printf("Промокод %s, скидка: %.2f\n", code, full-total)
		fmt.Printf("Итого к оплате: %.2f\n", total)
	}
}
