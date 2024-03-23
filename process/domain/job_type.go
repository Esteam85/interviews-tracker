package domain

import (
	"fmt"
	"strings"
)

type JobType int

const (
	contract JobType = iota
	fulltime
)

var jobTypeMap = map[string]JobType{
	"contract": contract,
	"fulltime": fulltime,
}

func ParseJobType(s string) (JobType, error) {
	if p, ok := jobTypeMap[strings.ToLower(s)]; ok {
		return p, nil
	}
	return 0, fmt.Errorf("invalid job type value: %q", s)
}
