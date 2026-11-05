package discount

import (
	"fmt"
	"sort"
	"strings"
)

// promoCodes — зарегистрированные промокоды: название → фабрика модификатора.
// Хранится фабрика, а не сам модификатор, потому что у модификатора
// WithLimitedDiscount есть счётчик, и для каждого расчёта нужен новый.
var promoCodes = make(map[string]func() PriceModifier)

// Register добавляет промокод code. factory вызывается при каждом
// обращении к ByCode и должна возвращать новый модификатор.
// Если такой промокод уже есть, Register вызывает panic.
func Register(code string, factory func() PriceModifier) {
	code = strings.ToUpper(code)
	if _, exists := promoCodes[code]; exists {
		panic("discount: промокод " + code + " зарегистрирован дважды")
	}
	promoCodes[code] = factory
}

// ByCode возвращает новый модификатор для промокода code (регистр букв не
// важен) или ошибку, если такого промокода нет.
func ByCode(code string) (PriceModifier, error) {
	factory, ok := promoCodes[strings.ToUpper(code)]
	if !ok {
		return nil, fmt.Errorf("промокод %q не найден", code)
	}
	return factory(), nil
}

// Codes возвращает названия всех зарегистрированных промокодов по алфавиту.
func Codes() []string {
	codes := make([]string, 0, len(promoCodes))
	for code := range promoCodes {
		codes = append(codes, code)
	}
	sort.Strings(codes)
	return codes
}
