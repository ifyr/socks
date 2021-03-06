package compiler

import (
	"errors"
	"strings"
)

type SCompiler struct {
	matchs map[string][]SMatch
}

func NewSCompiler() *SCompiler {
	return &SCompiler{
		matchs: make(map[string][]SMatch),
	}
}

func (this *SCompiler) Add(host string, rule []string) error {
	host = strings.ToLower(host)

	for i := 0; i < len(rule); i++ {

		smatch, err := NewSMatch(rule[i])

		if err != nil {
			return err
		}

		this.matchs[host] = append(this.matchs[strings.ToLower(host)], smatch)
	}

	return nil
}

func (this *SCompiler) Replace(host string, src string) (dst string, err error) {
	host = strings.ToLower(host)

	var exist bool
	var rules []SMatch

	if rules, exist = this.matchs["."]; false == exist {
		for i := len(host); i > 0; i = strings.LastIndexByte(host[:i], '.') {

			rules, exist = this.matchs[host[i:]]

			if true == exist {
				break
			}
		}
	}

	if 0 == len(rules) {
		rules = this.matchs[host]
	}

	for _, match := range rules {
		if dst, err := match.Replace(src); err == nil {
			return dst, nil
		}
	}

	return src, errors.New("regular expression does not match")
}
