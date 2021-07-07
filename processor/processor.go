package processor

import (
	"errors"
	"regexp"
	"strings"
)

const (
	containsLetterPattern = ".*[a-zA-Z]"
)

type Occurence struct {
	Word  string
	Count int64
}

func (s *Service) Process(limit int) error {
	words := s.source.Parse()
	if words == nil {
		return errors.New("could not parse source")
	}

	for w := range words {
		w = strings.ToLower(w)
		if !containsLetter(w) {
			continue
		}

		err := s.storage.ProcessWord(strings.Trim(w, "-'"))
		if err != nil {
			return err
		}
	}

	occurences := s.storage.GetTopFrequentOccurences(limit)
	for _, o := range occurences {
		err := s.output.Write(o)
		if err != nil {
			return err
		}
	}

	return nil
}

func containsLetter(s string) bool {
	matched, err := regexp.MatchString(containsLetterPattern, s)
	return matched && err == nil
}
