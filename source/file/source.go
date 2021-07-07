package file

import (
	"bufio"
	"os"
	"regexp"
	"strings"

	"github.com/crashage/wordcounter/utils"
)

const (
	isHtmlFilePattern = "[*]?.html"
	isTxtFilePattern  = "[*]?.txt"
)

type Source struct {
	f        *os.File
	isParsed bool
	out      chan string
}

type scanner interface {
	Scan() bool
	Text() string
}

func NewSource(name string) (*Source, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}

	return &Source{
		f:   f,
		out: make(chan string),
	}, nil
}

func (s *Source) Parse() <-chan string {
	if s.isParsed {
		return s.out
	}

	var sc scanner

	if isTxtFile(s.f.Name()) {
		sc = bufio.NewScanner(s.f)
	} else if isHtmlFile(s.f.Name()) {
		sc = newHtmlScanner(s.f)
	} else {
		return nil
	}

	go func(sc scanner, out chan string) {
		for sc.Scan() {
			for _, w := range strings.FieldsFunc(sc.Text(), utils.WordSplit) {
				out <- w
			}
		}
		close(out)
	}(sc, s.out)

	s.isParsed = true
	return s.out
}

func isHtmlFile(s string) bool {
	matched, err := regexp.MatchString(isHtmlFilePattern, s)
	return matched && err == nil
}

func isTxtFile(s string) bool {
	matched, err := regexp.MatchString(isTxtFilePattern, s)
	return matched && err == nil
}
