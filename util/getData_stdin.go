package util

import (
	"flag"
	"os"
)

func GetData_stdin() (*os.File, File, error) {
	// stdin, err := io.ReadAll(os.Stdin)
	var f File
	stdin := os.Stdin
	fi, err := stdin.Stat()

	if err != nil {
		return nil, File{}, err
	}

	// assign filedata to File
	if flag.NFlag() == 0 {
		f = NewFile(fi.Name())
	} else {
		f = File{
			Filepath: fi.Name(),
			Option: map[string]bool{
				"bytes": bytesOpt,
				"lines": linesOpt,
				"words": wordsOpt,
				"chars": charsOpt,
			},
		}
	}

	return stdin, f, nil
}
