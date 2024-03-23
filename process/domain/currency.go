package domain

import (
	"fmt"
	"strings"
)

type Currency int

const (
	usd Currency = iota
	clp
)

var currencyMap = map[string]Currency{
	"usd": usd,
	"clp": clp,
}

func ParseCurrency(s string) (Currency, error) {
	if c, ok := currencyMap[strings.ToLower(s)]; ok {
		return c, nil
	}
	return 0, fmt.Errorf("invalid currency value: %q", s)
}
