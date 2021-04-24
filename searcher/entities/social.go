package entities

import (
	"fmt"
	"regexp"
)

type SearchSocials struct {
	vkReg       *regexp.Regexp
	facebookReg *regexp.Regexp
}

func NewSearchSocial() (*SearchSocials, error) {

	reVK, err := regexp.Compile(`(https{0,1}://)?(www\.)?(vk.com/)(id\d|[a-zA-z][a-zA-Z0-9_.]{2,})`)
	if err != nil {
		return nil, fmt.Errorf("can't create regexp for vk: %w", err)
	}
	reFacebook, err := regexp.Compile(`(?:(?:http|https)://)?(?:www.)?facebook.com/(?:(?:\w)*#!/)?(?:pages/)?(?:[?\w\-]*/)?(?:profile.php\?id=(\d*))?([\w\-]*)?`)
	if err != nil {
		return nil, fmt.Errorf("can't create regexp for facebook: %w", err)
	}

	return &SearchSocials{
		vkReg:       reVK,
		facebookReg: reFacebook,
	}, nil
}

func (s *SearchSocials) Search(data string) ([]string, error) {

	socials := make([]string, 0)

	vks := s.vkReg.FindAllString(data, -1)
	fbs := s.facebookReg.FindAllString(data, -1)

	socials = append(vks, fbs...)

	return socials, nil
}

func (s *SearchSocials) Title() string {
	return "Social network"
}
