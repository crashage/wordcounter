package file

import (
	"fmt"
	"os"

	"github.com/crashage/wordcounter/processor"
)

type Output struct {
	f *os.File
}

func NewOutput(name string) (*Output, error) {
	f, err := os.Create(name)
	if err != nil {
		return nil, err
	}

	return &Output{
		f: f,
	}, nil
}

func (o *Output) Write(oc processor.Occurence) error {
	_, err := o.f.WriteString(fmt.Sprintf("%s : %d\n", oc.Word, oc.Count))
	return err
}
