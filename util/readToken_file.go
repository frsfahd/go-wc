package util

import (
	"bufio"
	"os"
)

func readLines(f *os.File) (countLines int) {
	// reset the pointer
	f.Seek(0, 0)

	scanner := bufio.NewScanner(f)
	// split by line
	scanner.Split(bufio.ScanLines)
	// count lines
	for scanner.Scan() {
		countLines++
	}

	return
}

func readBytes(f *os.File) (countBytes int) {
	// reset the pointer
	f.Seek(0, 0)

	scanner := bufio.NewScanner(f)
	// split by bytes
	scanner.Split(bufio.ScanBytes)
	// count bytes
	for scanner.Scan() {
		countBytes++
	}

	return
}

func readWords(f *os.File) (countWords int) {
	// reset the pointer
	f.Seek(0, 0)

	scanner := bufio.NewScanner(f)
	// split by words
	scanner.Split(bufio.ScanWords)
	// count words
	for scanner.Scan() {
		countWords++
	}

	return
}

func readChars(f *os.File) (countChars int) {
	// reset the pointer
	f.Seek(0, 0)

	scanner := bufio.NewScanner(f)
	// split by runes
	scanner.Split(bufio.ScanRunes)
	// count runes
	for scanner.Scan() {
		countChars++
	}

	return
}
