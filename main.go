package main

import (
	"flag"
	"log"

	ofile "github.com/crashage/wordcounter/output/file"
	"github.com/crashage/wordcounter/output/stdout"
	"github.com/crashage/wordcounter/processor"
	ifile "github.com/crashage/wordcounter/source/file"
	"github.com/crashage/wordcounter/source/stdin"
	iurl "github.com/crashage/wordcounter/source/url"
	"github.com/crashage/wordcounter/storage/memory"
)

var (
	url    string
	text   string
	file   string
	output string
	limit  int
)

func main() {
	parseFlags()

	source := in()
	output := out()
	storage := memory.NewStorage()

	processor := processor.NewService(storage, source, output)
	err := processor.Process(limit)
	if err != nil {
		log.Fatalln(err)
	}
}

func parseFlags() {
	flag.StringVar(&url, "u", "", "Url input source, has the lowest priority if multiple input sources are provided.")
	flag.StringVar(&text, "t", "", "Text input source, has the highest priority if multiple input sources are provided.")
	flag.StringVar(&file, "f", "", "File input source, has the second highest priority if multiple input sources are provided.")
	flag.StringVar(&output, "o", "", "Name of output file. Keep blank to use stdout.")
	flag.IntVar(&limit, "l", -1, "Limit of top frequent words printed to output. Keep blank to print all words.")

	flag.Parse()
}

func in() processor.Source {
	if len(text) > 0 {
		return stdin.NewSource(text)
	}

	if len(file) > 0 {
		s, err := ifile.NewSource(file)
		if err != nil {
			log.Fatalln(err)
		}
		return s
	}

	if len(url) > 0 {
		s, err := iurl.NewSource(url)
		if err != nil {
			log.Fatalln(err)
		}
		return s
	}

	log.Fatalln("no input provided")

	return nil
}

func out() processor.Output {
	if len(output) == 0 {
		return stdout.NewOutput()
	}

	o, err := ofile.NewOutput(output)
	if err != nil {
		log.Fatalln(err)
	}

	return o
}
