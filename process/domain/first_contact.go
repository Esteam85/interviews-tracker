package domain

import (
	"fmt"
	"strings"
	"time"
)

type RecruitmentChannel int

const (
	mail RecruitmentChannel = iota
	linkedinChat
	whatsapp
	phoneCall
	textMessage
)

var recruitmentChannelMap = map[string]RecruitmentChannel{
	"mail":         mail,
	"linkedinchat": linkedinChat,
	"whatsapp":     whatsapp,
	"phonecall":    phoneCall,
	"textmessage":  textMessage,
}

func ParseRecruitmentChannel(s string) (RecruitmentChannel, error) {
	if r, ok := recruitmentChannelMap[strings.ToLower(s)]; ok {
		return r, nil
	}
	return 0, fmt.Errorf("invalid recruitment channel value: %q", s)

}

type FirstContact struct {
	Ate          time.Time          `json:"date"`
	Channel      RecruitmentChannel `json:"channel,omitempty"`
	AnsweredDate time.Time          `json:"answeredDate"`
}

func NewFirstContact(date string, channel string, options ...func(f *FirstContact) error) (*FirstContact, error) {
	fCDate, err := time.Parse("2006-01-02", date)
	if err != nil {
		return &FirstContact{}, err
	}
	c, err := ParseRecruitmentChannel(channel)
	if err != nil {
		return &FirstContact{}, err
	}

	f := &FirstContact{
		Ate:          fCDate,
		Channel:      c,
		AnsweredDate: time.Time{},
	}
	for _, o := range options {
		err = o(f)
		if err != nil {
			return &FirstContact{}, err
		}
	}
	return f, nil
}

func WithAnsweredDate(s string) func(contact *FirstContact) error {
	return func(f *FirstContact) error {
		a, err := time.Parse("2006-01-02", s)
		if err != nil {
			return err
		}

		if f.Ate.After(a) {
			return fmt.Errorf("invalid answered date %s, it can't be setted before %s", s, f.Ate.String())
		}

		f.AnsweredDate = a
		return nil
	}
}
