package util

import (
	"bufio"
	"io"
)

func readLines_stdin(r io.Reader) (countLines int) {

	scanner := bufio.NewScanner(r)
	// split by line
	scanner.Split(bufio.ScanLines)
	// count lines
	for scanner.Scan() {
		countLines++
	}

	return
}

func readBytes_stdin(r io.Reader) (countBytes int) {

	scanner := bufio.NewScanner(r)
	// split by bytes
	scanner.Split(bufio.ScanBytes)
	// count bytes
	for scanner.Scan() {
		countBytes++
	}

	return
}

func readWords_stdin(r io.Reader) (countWords int) {

	scanner := bufio.NewScanner(r)
	// split by words
	scanner.Split(bufio.ScanWords)
	// count words
	for scanner.Scan() {
		countWords++
	}

	return
}

func readChars_stdin(r io.Reader) (countChars int) {

	scanner := bufio.NewScanner(r)
	// split by runes
	scanner.Split(bufio.ScanRunes)
	// count runes
	for scanner.Scan() {
		countChars++
	}

	return
}
