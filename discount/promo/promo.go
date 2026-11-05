// Package promo регистрирует промокоды магазина в пакете discount.
//
// У пакета нет экспортируемых имён: всю работу делает init. Поэтому его
// подключают пустым импортом (blank import):
//
//	import _ "laba7/discount/promo"
package promo

import "laba7/discount"

func init() {
	// −50 % на первые две единицы товара в корзине.
	discount.Register("SALE50", func() discount.PriceModifier {
		return discount.WithLimitedDiscount(50, 2)
	})
	// −30 % на первую единицу товара в корзине.
	discount.Register("FIRST30", func() discount.PriceModifier {
		return discount.WithLimitedDiscount(30, 1)
	})
	// −10 % на все товары, но цена не ниже 1 рубля.
	discount.Register("MINUS10", func() discount.PriceModifier {
		return discount.CombineModifiers(
			discount.WithDiscount(10),
			discount.WithMinimumPrice(1),
		)
	})
}
