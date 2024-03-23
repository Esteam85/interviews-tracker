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
	"linkedin":       linkedin,
	"company_portal": companyPortal,
	"computrabajo":   compuTrabajo,
	"get_onboard":    getOnBoard,
}
var invertPlatformMap = map[Platform]string{
	linkedin:      "linkedin",
	companyPortal: "company_portal",
	compuTrabajo:  "computrabajo",
	getOnBoard:    "get_onboard",
}

func ParsePlatform(s string) (Platform, error) {
	if p, ok := platformMap[strings.ToLower(s)]; ok {
		return p, nil
	}
	return 0, fmt.Errorf("invalid platform type value: %q", s)
}

func (p Platform) String() string {
	return invertPlatformMap[p]
}
