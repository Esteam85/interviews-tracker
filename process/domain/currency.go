package domain

import (
	"errors"
	"fmt"
	"strings"
)

var ErrInvalidCurrency = errors.New("invalid currency")

type Currency int

const (
	usd Currency = iota
	clp
)

var currencyMap = map[string]Currency{
	"usd": usd,
	"clp": clp,
}

var invertCurrencyMap = map[Currency]string{
	usd: "usd",
	clp: "clp",
}

func ParseCurrency(s string) (Currency, error) {
	if c, ok := currencyMap[strings.ToLower(s)]; ok {
		return c, nil
	}
	return 0, fmt.Errorf("%w with value: %s", ErrInvalidCurrency, s)
}

func (c Currency) String() string {
	return invertCurrencyMap[c]
}
