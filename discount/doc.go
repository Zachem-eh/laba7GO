// Package discount содержит модификаторы цены и промокоды.
//
// Модификатор цены (PriceModifier) — это функция, которая принимает цену и
// возвращает новую. Функции-фабрики WithDiscount, WithTax, WithRounding,
// WithMinimumPrice и WithLimitedDiscount создают модификаторы с помощью
// замыканий, а CombineModifiers объединяет несколько модификаторов в один.
//
// Промокоды регистрируются функцией Register (обычно в init другого пакета,
// см. discount/promo) и выдаются функцией ByCode.
package discount
