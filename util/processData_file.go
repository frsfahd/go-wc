package util

import (
	"fmt"
	"os"
)

func ProcessData_file(f File) string {

	// open file
	txtFile, err := os.Open(f.Filepath)

	if err != nil {
		ExitGracefully(err)
	}

	// count token based on params/flags
	if f.Option["lines"] {
		countResult = countResult + fmt.Sprintf("%d ", readLines(txtFile))
	}
	if f.Option["bytes"] {
		countResult = countResult + fmt.Sprintf("%d ", readBytes(txtFile))
	}
	if f.Option["words"] {
		countResult = countResult + fmt.Sprintf("%d ", readWords(txtFile))
	}
	if f.Option["chars"] {
		countResult = countResult + fmt.Sprintf("%d ", readChars(txtFile))
	}

	// close file
	defer txtFile.Close()

	return countResult + " " + f.Filepath

}
