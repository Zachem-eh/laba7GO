package receipt

import (
	"fmt"
	"laba7/cart"
	"os"
	"text/tabwriter"
)

func PrintBasket() {
	fmt.Println("=== Чек ===")
	var cartTotal float64
	myBasket := cart.GetBasket()
	if len(myBasket) == 0 {
		fmt.Println("Корзина пуста.")
		return
	} else {
		// Параметры: os.Stdout (куда писать), 0 (minwidth), 0 (tabwidth), 3 (padding), ' ' (padchar), 0 (flags)
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', 0)
		fmt.Fprintln(w, "ID\tНАЗВАНИЕ\tЦЕНА\t")
		fmt.Fprintln(w, "--\t--------\t-----\t")

		for i := 0; i < len(myBasket); i++ {
			fmt.Fprintf(w, "%d\t%s\t%.2f\t\n", (myBasket)[i].ID, (myBasket)[i].Name, (myBasket)[i].Price)
			cartTotal += (myBasket)[i].Price
		}
		fmt.Fprintf(w, "Итого: %.2f\n", cartTotal)
		w.Flush()
	}
}
