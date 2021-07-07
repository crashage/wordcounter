package stdin

import (
	"strings"

	"github.com/crashage/wordcounter/utils"
)

type Source struct {
	raw      string
	isParsed bool
	out      chan string
}

func NewSource(data string) *Source {
	return &Source{
		raw:      data,
		isParsed: false,
		out:      make(chan string),
	}
}

func (s *Source) Parse() <-chan string {
	if s.isParsed {
		return s.out
	}

	go func(raw string, out chan string) {
		for _, w := range strings.FieldsFunc(raw, utils.WordSplit) {
			out <- w
		}

		close(out)
	}(s.raw, s.out)

	s.isParsed = true

	return s.out
}
