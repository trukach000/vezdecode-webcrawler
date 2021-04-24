package entities

import (
	"fmt"
	"regexp"
)

type SearchEmail struct {
	emailReg *regexp.Regexp
}

func NewSearchEmail() (*SearchEmail, error) {

	emailReg, err := regexp.Compile("(?i)([A-ZЁА-Я0-9._%+-]+@[ЁА-ЯA-Z0-9.-]+\\.[ЁА-ЯA-Z]{2,24})")
	if err != nil {
		return nil, fmt.Errorf("can't create regexp for email: %w", err)
	}

	return &SearchEmail{
		emailReg: emailReg,
	}, nil
}

func (s *SearchEmail) Search(data string) ([]string, error) {

	emails := s.emailReg.FindAllString(data, -1)

	return emails, nil
}

func (s *SearchEmail) Title() string {
	return "Email"
}
