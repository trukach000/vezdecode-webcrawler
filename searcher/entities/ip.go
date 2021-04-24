package entities

import (
	"fmt"
	"regexp"
)

type SearchIp struct {
	IPv4Reg *regexp.Regexp
	IPv6Reg *regexp.Regexp
}

func NewSearchIP() (*SearchIp, error) {

	reIPv4, err := regexp.Compile(`(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)(\.(25[0-5]|2[0-4][0-9]|[01]?[0-9][0-9]?)){3}`)
	if err != nil {
		return nil, fmt.Errorf("can't create regexp for IP v4: %w", err)
	}
	reIPv6, err := regexp.Compile(`(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`)
	if err != nil {
		return nil, fmt.Errorf("can't create regexp for IP v6: %w", err)
	}

	return &SearchIp{
		IPv4Reg: reIPv4,
		IPv6Reg: reIPv6,
	}, nil
}

func (s *SearchIp) Search(data string) ([]string, error) {

	ips := make([]string, 0)

	ip4Matches := s.IPv4Reg.FindAllString(data, -1)
	ip6Matches := s.IPv6Reg.FindAllString(data, -1)

	ips = append(ip4Matches, ip6Matches...)

	return ips, nil
}

func (s *SearchIp) Title() string {
	return "IP address"
}
