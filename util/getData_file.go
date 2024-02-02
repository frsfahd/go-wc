package util

import (
	"flag"
)

func GetData_file() (File, error) {
	var f File

	// assign filedata to File
	if flag.NFlag() == 0 {
		f = NewFile(flag.Arg(0))
	} else {
		f = File{
			Filepath: flag.Arg(0),
			Option: map[string]bool{
				"bytes": bytesOpt,
				"lines": linesOpt,
				"words": wordsOpt,
				"chars": charsOpt,
			},
		}
	}

	// debug
	// fmt.Printf("num of flag: %d\n", flag.NFlag())
	// fmt.Printf("num of args: %d\n", flag.NArg())
	// fmt.Printf("all args (%d): %v\n", len(os.Args), os.Args)
	// fmt.Printf("all flags: %v", flag.CommandLine)

	return f, nil

}
