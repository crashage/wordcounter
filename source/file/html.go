package file

import (
	"os"

	"github.com/crashage/wordcounter/utils"
	"golang.org/x/net/html"
)

type htmlScanner struct {
	t       *html.Tokenizer
	prevTag string
}

func newHtmlScanner(f *os.File) *htmlScanner {
	return &htmlScanner{
		t: html.NewTokenizer(f),
	}
}

func (s *htmlScanner) Scan() bool {
	for {
		tt := s.t.Next()
		switch tt {
		case html.ErrorToken:
			return false
		case html.StartTagToken:
			s.prevTag = s.t.Token().Data
		case html.TextToken:
			return true
		}
	}
}

func (s *htmlScanner) Text() string {
	if utils.SkipTag(s.prevTag) {
		return ""
	}
	return string(s.t.Text())
}
