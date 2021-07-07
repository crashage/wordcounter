package processor

type Storage interface {
	ProcessWord(word string) error
	GetTopFrequentOccurences(count int) []Occurence
}

type Source interface {
	Parse() <-chan string
}

type Output interface {
	Write(Occurence) error
}

type Service struct {
	storage Storage
	source  Source
	output  Output
}

func NewService(storage Storage, source Source, output Output) *Service {
	return &Service{
		storage: storage,
		source:  source,
		output:  output,
	}
}
