package memory

import (
	"github.com/wangjia184/sortedset"

	"github.com/crashage/wordcounter/processor"
)

type Storage struct {
	set *sortedset.SortedSet
}

func NewStorage() *Storage {
	return &Storage{
		set: sortedset.New(),
	}
}

func (s *Storage) ProcessWord(word string) error {
	newScore := int64(1)
	node := s.set.GetByKey(word)

	if node != nil {
		newScore += int64(node.Score())
	}

	s.set.AddOrUpdate(word, sortedset.SCORE(newScore), struct{}{})

	return nil
}

func (s *Storage) GetTopFrequentOccurences(count int) []processor.Occurence {
	vals := s.set.GetByRankRange(-1, -count, false)
	res := make([]processor.Occurence, 0, len(vals))

	for _, w := range vals {
		res = append(res, processor.Occurence{
			Word:  w.Key(),
			Count: int64(w.Score()),
		})
	}

	return res
}
