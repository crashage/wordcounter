package stdout

import (
	"fmt"

	"github.com/crashage/wordcounter/processor"
)

type Output struct{}

func NewOutput() *Output {
	return &Output{}
}

func (o *Output) Write(oc processor.Occurence) error {
	fmt.Printf("%s : %d\n", oc.Word, oc.Count)
	return nil
}
