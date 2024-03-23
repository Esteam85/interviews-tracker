package domain

import (
	"fmt"
	"strings"
)

type Platform int

const (
	linkedin Platform = iota
	companyPortal
	getOnBoard
	compuTrabajo
)

var platformMap = map[string]Platform{
	"linkedin":      linkedin,
	"companyportal": companyPortal,
	"computrabajo":  compuTrabajo,
	"getOnBoard":    getOnBoard,
}

func ParsePlatform(s string) (Platform, error) {
	if p, ok := platformMap[strings.ToLower(s)]; ok {
		return p, nil
	}
	return 0, fmt.Errorf("invalid platform type value: %q", s)
}
