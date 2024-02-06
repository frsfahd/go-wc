package util

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func ProcessData_stdin(txtFile *os.File, f File) (string, error) {

	// Read all data from stdin
	data, err := io.ReadAll(txtFile)
	if err != nil {
		return "", err
	}

	// Convert data to string
	s := string(data)

	// Create new readers for each function
	readerLines := strings.NewReader(s)
	readerBytes := strings.NewReader(s)
	readerWords := strings.NewReader(s)
	readerChars := strings.NewReader(s)

	// count token based on params/flags
	if f.Option["lines"] {
		countResult = countResult + fmt.Sprintf("%d ", readLines_stdin(readerLines))
	}
	if f.Option["bytes"] {
		countResult = countResult + fmt.Sprintf("%d ", readBytes_stdin(readerBytes))
	}
	if f.Option["words"] {
		countResult = countResult + fmt.Sprintf("%d ", readWords_stdin(readerWords))
	}
	if f.Option["chars"] {
		countResult = countResult + fmt.Sprintf("%d ", readChars_stdin(readerChars))
	}

	return countResult + " ", nil

}
