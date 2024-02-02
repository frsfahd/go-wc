package util

import (
	"flag"
	"fmt"
	"os"
)

var (
	bytesOpt bool
	linesOpt bool
	wordsOpt bool
	charsOpt bool
)

var countResult string

type File struct {
	Filepath string
	Option   map[string]bool
}

func NewFile(f string) File {
	return File{
		Filepath: f,
		Option: map[string]bool{
			"bytes": true,  // -c
			"lines": true,  // -l
			"words": true,  // -w
			"chars": false, // -m
		},
	}
}

func init() {
	// define flags
	flag.BoolVar(&bytesOpt, "c", false, "print the byte counts")      // -c
	flag.BoolVar(&linesOpt, "l", false, "print the newline counts")   // -l
	flag.BoolVar(&wordsOpt, "w", false, "print the word counts")      // -w
	flag.BoolVar(&charsOpt, "m", false, "print the character counts") // -m

	flag.Parse()
}

func ExitGracefully(err error) {
	fmt.Fprintf(os.Stderr, "error: %v", err)
	os.Exit(1)
}
