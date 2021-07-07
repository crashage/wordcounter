package url

import (
	"io/ioutil"
	"net/http"
	"strings"

	"golang.org/x/net/html"

	"github.com/crashage/wordcounter/utils"
)

type bodyType int

const (
	txtType bodyType = iota
	htmlType
)

type Source struct {
	raw      string
	isParsed bool
	out      chan string
	typ      bodyType
}

func NewSource(url string) (*Source, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return &Source{
		raw: string(raw),
		out: make(chan string),
		typ: getContentType(resp),
	}, nil
}

func (s *Source) Parse() <-chan string {
	if s.isParsed {
		return s.out
	}

	go func(typ bodyType, raw string, out chan string) {
		var words []string
		switch typ {
		case txtType:
			words = strings.FieldsFunc(raw, utils.WordSplit)
		case htmlType:
			words = parseHtml(raw)
		}

		for _, w := range words {
			out <- w
		}

		close(out)
	}(s.typ, s.raw, s.out)

	s.isParsed = true
	return s.out
}

func getContentType(r *http.Response) bodyType {
	cts := strings.Split(r.Header.Get("Content-Type"), ";")

	for _, ct := range cts {
		if ct == "text/html" {
			return htmlType
		}
	}

	return txtType
}

func parseHtml(raw string) []string {
	tok := html.NewTokenizer(strings.NewReader(raw))
	res := make([]string, 0)
	prevTag := ""

	for {
		tt := tok.Next()
		switch tt {
		case html.ErrorToken:
			return res
		case html.StartTagToken:
			prevTag = tok.Token().Data
		case html.TextToken:
			if utils.SkipTag(prevTag) {
				continue
			}
			words := strings.FieldsFunc(string(tok.Text()), utils.WordSplit)
			res = append(res, words...)
		}
	}
}
