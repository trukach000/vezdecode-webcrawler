package entities

import (
	"fmt"
	"regexp"
)

type SearchPhone struct {
	mobileReg *regexp.Regexp
}

func NewSearchPhone() (*SearchPhone, error) {

	phoneReg, err := regexp.Compile("(?:^|[^\\d])((?:\\+7|8)[- _]*\\(?:?[- _]*(?:\\d{3}[- _]*\\)?(?:[- _]*\\d){7}|\\d\\d[- _]*\\d\\d[- _]*\\)?(?:[- _]*\\d){6}))")
	if err != nil {
		return nil, fmt.Errorf("can't create regexp for phone: %w", err)
	}

	return &SearchPhone{
		mobileReg: phoneReg,
	}, nil
}

func (s *SearchPhone) Search(data string) ([]string, error) {

	emailSubmatches := s.mobileReg.FindAllStringSubmatch(data, -1)

	emails := make([]string, 0)
	for _, es := range emailSubmatches {
		emails = append(emails, es[1])
	}

	return emails, nil
}

func (s *SearchPhone) Title() string {
	return "phone"
}
