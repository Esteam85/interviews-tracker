package domain

import (
	"fmt"
	"strings"
)

type PostulationType int

const (
	own PostulationType = iota
	recruiter
	reference
)

var postulationTypeMap = map[string]PostulationType{
	"own":       own,
	"recruiter": recruiter,
	"reference": reference,
}

var invertPostulationTypeMap = map[PostulationType]string{
	own:       "own",
	recruiter: "recruiter",
	reference: "reference",
}

func ParsePostulationType(s string) (PostulationType, error) {
	if p, ok := postulationTypeMap[strings.ToLower(s)]; ok {
		return p, nil
	}
	return 0, fmt.Errorf("invalid postulation type value: %q", s)
}

func (p PostulationType) String() string {
	return invertPostulationTypeMap[p]
}
