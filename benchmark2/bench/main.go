package bench

import (
	"strings"
)

func Cat(strslice []string) string {
	s := ""
	for _, v := range strslice {
		s += v
		s += " "
	}
	return s
}

func Join(strslice []string) string {
	return strings.Join(strslice, " ")
}
